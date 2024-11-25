package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"passgen/internal/clipboard"
	"passgen/internal/crypto"
	"passgen/internal/qr"
)

func main() {
	var length int
	var useBip39 bool
	var customText string
	var outputFile string
	var qrSize int
	var usePipe bool
	var useEncrypt bool
	var encryptText string
	var decrypt string
	var password string

	flag.IntVar(&length, "l", 0, "Password length (shorthand)")
	flag.BoolVar(&useBip39, "b", false, "Generate BIP39 mnemonic")
	flag.StringVar(&customText, "c", "", "Use custom text for QR code")
	flag.StringVar(&outputFile, "o", "", "Output file for QR code (PNG format)")
	flag.IntVar(&qrSize, "s", 256, "QR code size in pixels (default: 256)")
	flag.BoolVar(&useEncrypt, "e", false, "Encrypt input")
	flag.StringVar(&encryptText, "text", "", "Text to encrypt")
	flag.StringVar(&decrypt, "d", "", "Text/file to decrypt")
	flag.StringVar(&password, "p", "", "Password for encryption/decryption")
	flag.Parse()

	// Проверяем pipe ввод
	stat, _ := os.Stdin.Stat()
	usePipe = (stat.Mode() & os.ModeCharDevice) == 0

	var result string

	switch {
	case usePipe:
		// Читаем из pipe
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Failed to read from pipe: %v", err)
		}
		text := strings.TrimSpace(string(bytes))

		// Если указан флаг -e, шифруем текст из pipe
		if useEncrypt {
			if password == "" {
				log.Fatal("Password required for encryption (-p flag)")
			}
			xchacha, err := crypto.NewXChaCha(password)
			if err != nil {
				log.Fatalf("Failed to initialize encryption: %v", err)
			}
			encrypted, err := xchacha.Encrypt([]byte(text))
			if err != nil {
				log.Fatalf("Failed to encrypt: %v", err)
			}
			result = encrypted
		} else {
			result = text
		}
	case useEncrypt:
		if password == "" {
			log.Fatal("Password required for encryption (-p flag)")
		}

		// Используем аргументы командной строки как текст для шифрования
		text := flag.Arg(0)
		if text == "" && encryptText == "" {
			log.Fatal("No text provided for encryption")
		}
		if text == "" {
			text = encryptText
		}

		xchacha, err := crypto.NewXChaCha(password)
		if err != nil {
			log.Fatalf("Failed to initialize encryption: %v", err)
		}

		encrypted, err := xchacha.Encrypt([]byte(text))
		if err != nil {
			log.Fatalf("Failed to encrypt: %v", err)
		}

		result = encrypted

		// Сохраняем в файл если указан
		if outputFile != "" {
			if strings.HasSuffix(outputFile, ".png") {
				// Сохраняем как QR код
				if err := qr.GenerateToFile(encrypted, outputFile, qrSize); err != nil {
					log.Printf("Failed to save QR code: %v", err)
				} else {
					fmt.Printf("QR code saved to: %s\n", outputFile)
				}
			} else {
				// Сохраняем как текстовый/бинарный файл
				if err := os.WriteFile(outputFile, []byte(encrypted), 0644); err != nil {
					log.Printf("Failed to save encrypted file: %v", err)
				} else {
					fmt.Printf("Encrypted data saved to: %s\n", outputFile)
				}
			}
		}
	case decrypt != "":
		if password == "" {
			log.Fatal("Password required for decryption (-p flag)")
		}

		var ciphertext string
		if _, err := os.Stat(decrypt); err == nil {
			// Читаем файл с автоопределением формата
			data, err := qr.ReadFromFile(decrypt)
			if err != nil {
				log.Fatalf("Failed to read file: %v", err)
			}
			ciphertext = data
		} else {
			// Используем текст напрямую
			ciphertext = decrypt
		}

		xchacha, err := crypto.NewXChaCha(password)
		if err != nil {
			log.Fatalf("Failed to initialize decryption: %v", err)
		}

		decrypted, err := xchacha.Decrypt(ciphertext)
		if err != nil {
			log.Fatalf("Failed to decrypt: %v", err)
		}

		result = string(decrypted)
	case customText != "":
		result = customText
	case useBip39:
		// BIP39 логика
		// ...
	default:
		// Обычный пароль
		// ...
	}

	fmt.Println(result)

	if err := clipboard.Copy(result); err != nil {
		log.Printf("Failed to copy to clipboard: %v", err)
	} else {
		fmt.Println("Copied to clipboard")
	}

	if outputFile != "" {
		if err := qr.GenerateToFile(result, outputFile, qrSize); err != nil {
			log.Printf("Failed to save QR code: %v", err)
		} else {
			fmt.Printf("QR code saved to: %s\n", outputFile)
		}
	} else {
		qrCode, err := qr.Generate(result)
		if err != nil {
			log.Printf("Failed to create QR code: %v", err)
		} else {
			fmt.Println("\nQR code:")
			fmt.Println(qrCode)
		}
	}
}
