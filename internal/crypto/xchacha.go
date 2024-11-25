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
	Version    int    `json:"version"`    // версия формата
}

type XChaCha struct {
	aead     cipher.AEAD
	password string
}

// NewXChaCha создает новый экземпляр XChaCha20-Poly1305 с паролем
func NewXChaCha(password string) (*XChaCha, error) {
	if password == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}

	return &XChaCha{
		password: password,
	}, nil
}

// Encrypt шифрует данные и возвращает JSON в base64
func (x *XChaCha) Encrypt(plaintext []byte) (string, error) {
	// Генерируем соль
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}

	// Генерируем ключ
	key, err := scrypt.Key([]byte(x.password), salt, scryptN, scryptR, scryptP, keyLen)
	if err != nil {
		return "", fmt.Errorf("failed to derive key: %v", err)
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AEAD: %v", err)
	}

	// Генерируем новый nonce для каждого шифрования
	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %v", err)
	}

	// Шифруем
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)

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

	// Кодируем весь JSON в base64
	return base64.StdEncoding.EncodeToString(jsonData), nil
}

// Decrypt расшифровывает JSON из base64
func (x *XChaCha) Decrypt(encoded string) ([]byte, error) {
	// Декодируем base64 в JSON
	jsonData, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	// Парсим JSON
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

	// Генерируем ключ
	key, err := scrypt.Key([]byte(x.password), salt, scryptN, scryptR, scryptP, keyLen)
	if err != nil {
		return nil, fmt.Errorf("failed to derive key: %v", err)
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AEAD: %v", err)
	}

	// Расшифровываем
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %v", err)
	}

	return plaintext, nil
}
