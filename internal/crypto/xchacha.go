package crypto

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/scrypt"
)

const (
	// Параметры для scrypt
	scryptN = 32768
	scryptR = 8
	scryptP = 1
	keyLen  = chacha20poly1305.KeySize
)

// CryptoData структура для хранения всех параметров шифрования
type CryptoData struct {
	Salt       string `json:"salt"`       // base64 encoded
	Nonce      string `json:"nonce"`      // base64 encoded
	Ciphertext string `json:"ciphertext"` // base64 encoded
	Version    int    `json:"version"`    // версия формата для обратной совместимости
}

type XChaCha struct {
	aead     cipher.AEAD
	password string
	salt     []byte
}

// NewXChaCha создает новый экземпляр XChaCha20-Poly1305 с паролем
func NewXChaCha(password string) (*XChaCha, error) {
	return &XChaCha{
		password: password,
	}, nil
}

func (x *XChaCha) initAEAD(salt []byte) error {
	key, err := scrypt.Key([]byte(x.password), salt, scryptN, scryptR, scryptP, keyLen)
	if err != nil {
		return fmt.Errorf("failed to derive key: %v", err)
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return fmt.Errorf("failed to create AEAD: %v", err)
	}

	x.aead = aead
	x.salt = salt
	return nil
}

// Encrypt шифрует данные и возвращает base64-encoded строку
func (x *XChaCha) Encrypt(plaintext []byte) (string, error) {
	// Генерируем соль
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}

	if err := x.initAEAD(salt); err != nil {
		return "", err
	}

	// Генерируем nonce
	nonce := make([]byte, x.aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %v", err)
	}

	// Шифруем
	ciphertext := x.aead.Seal(nil, nonce, plaintext, nil)

	// Создаем структуру с данными
	data := CryptoData{
		Salt:       base64.StdEncoding.EncodeToString(salt),
		Nonce:      base64.StdEncoding.EncodeToString(nonce),
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
		Version:    1,
	}

	// Сериализуем в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %v", err)
	}

	return base64.StdEncoding.EncodeToString(jsonData), nil
}

// Decrypt расшифровывает base64-encoded строку
func (x *XChaCha) Decrypt(encoded string) ([]byte, error) {
	// Декодируем JSON
	jsonData, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	var data CryptoData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %v", err)
	}

	// Проверяем версию
	if data.Version != 1 {
		return nil, fmt.Errorf("unsupported version: %d", data.Version)
	}

	// Декодируем компоненты
	salt, err := base64.StdEncoding.DecodeString(data.Salt)
	if err != nil {
		return nil, fmt.Errorf("failed to decode salt: %v", err)
	}

	nonce, err := base64.StdEncoding.DecodeString(data.Nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to decode nonce: %v", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data.Ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ciphertext: %v", err)
	}

	// Инициализируем AEAD с солью
	if err := x.initAEAD(salt); err != nil {
		return nil, err
	}

	// Расшифровываем
	plaintext, err := x.aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %v", err)
	}

	return plaintext, nil
}
