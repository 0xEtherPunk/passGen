package bip39

import (
	"crypto/sha256"
	"errors"
	"os"
	"strings"
)

const (
	// Длина энтропии в битах
	ENT128 = 128
	ENT160 = 160
	ENT192 = 192
	ENT224 = 224
	ENT256 = 256
)

type Mnemonic struct {
	language string
	wordlist []string
}

// New создает новый генератор BIP39 мнемоники
func New(language string) (*Mnemonic, error) {
	if language == "" {
		language = "en"
	}

	wordlist, err := LoadWordlist(language)
	if err != nil {
		return nil, err
	}

	return &Mnemonic{
		language: language,
		wordlist: wordlist,
	}, nil
}

// Generate создает новую мнемоническую фразу с указанной энтропией
func (m *Mnemonic) Generate(strength int) (string, error) {
	if strength != ENT128 && strength != ENT160 &&
		strength != ENT192 && strength != ENT224 &&
		strength != ENT256 {
		return "", errors.New("invalid entropy length")
	}

	// Получаем энтропию из /dev/urandom
	f, err := os.Open("/dev/urandom")
	if err != nil {
		return "", err
	}
	defer f.Close()

	entropy := make([]byte, strength/8)
	if _, err := f.Read(entropy); err != nil {
		return "", err
	}

	return m.EntropyToMnemonic(entropy)
}

// EntropyToMnemonic конвертирует байты энтропии в мнемоническую фразу
func (m *Mnemonic) EntropyToMnemonic(entropy []byte) (string, error) {
	if len(entropy)*8 < ENT128 || len(entropy)*8 > ENT256 {
		return "", errors.New("invalid entropy length")
	}

	// Создаем контрольную сумму
	hash := sha256.Sum256(entropy)
	checksumBits := len(entropy) * 8 / 32
	checksum := hash[0] >> (8 - checksumBits)

	// Комбинируем энтропию и контрольную сумму
	combined := make([]byte, len(entropy)+1)
	copy(combined, entropy)
	combined[len(entropy)] = checksum

	// Конвертируем биты в индексы слов
	var words []string
	for i := 0; i < len(combined)*8/11; i++ {
		// Извлекаем 11 бит для индекса
		index := extractBits(combined, i*11, 11)
		if index >= len(m.wordlist) {
			return "", errors.New("invalid word index")
		}
		words = append(words, m.wordlist[index])
	}

	// Соединяем слова с правильным разделителем
	delimiter := " "
	if m.language == "ja" {
		delimiter = "\u3000" // Японский использует идеографический пробел
	}

	return strings.Join(words, delimiter), nil
}

// extractBits извлекает указанное количество бит из байтового массива
func extractBits(data []byte, start, length int) int {
	var result int

	for i := 0; i < length; i++ {
		byteIndex := (start + i) / 8
		bitIndex := 7 - ((start + i) % 8)

		if byteIndex >= len(data) {
			break
		}

		if data[byteIndex]&(1<<bitIndex) != 0 {
			result |= 1 << (length - 1 - i)
		}
	}

	return result
}

// Check проверяет валидность мнемонической фразы
func (m *Mnemonic) Check(mnemonic string) bool {
	words := strings.Fields(mnemonic)

	// Проверяем количество слов
	if len(words) != 12 && len(words) != 15 &&
		len(words) != 18 && len(words) != 21 &&
		len(words) != 24 {
		return false
	}

	// Проверяем каждое слово в словаре
	for _, word := range words {
		found := false
		for _, dictWord := range m.wordlist {
			if word == dictWord {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// TODO: добавить проверку контрольной суммы
	return true
}
