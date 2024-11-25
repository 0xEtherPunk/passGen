package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"passgen/internal/bip39"
	"passgen/internal/clipboard"
	"passgen/internal/crypto"
	"passgen/internal/generator"
	"passgen/internal/qr"
)

func main() {
	var (
		length      int
		useBip39    bool
		customText  string
		outputFile  string
		qrSize      int
		usePipe     bool
		useEncrypt  bool
		encryptText string
		decrypt     string
		password    string
	)

	// Флаги для языков BIP39
	var langEn, langRu, langJp, langCn, langFr, langIt, langKo, langEs bool

	flag.IntVar(&length, "l", 0, "Password length (shorthand)")
	flag.BoolVar(&useBip39, "b", false, "Generate BIP39 mnemonic")
	flag.StringVar(&customText, "c", "", "Use custom text for QR code")
	flag.StringVar(&outputFile, "o", "", "Output file for QR code (PNG format)")
	flag.IntVar(&qrSize, "s", 256, "QR code size in pixels (default: 256)")
	flag.BoolVar(&useEncrypt, "e", false, "Encrypt input")
	flag.StringVar(&encryptText, "text", "", "Text to encrypt")
	flag.StringVar(&decrypt, "d", "", "Text/file to decrypt")
	flag.StringVar(&password, "p", "", "Password for encryption/decryption")

	// Добавляем флаги языков
	flag.BoolVar(&langEn, "en", false, "Use English wordlist")
	flag.BoolVar(&langRu, "ru", false, "Use Russian wordlist")
	flag.BoolVar(&langJp, "jp", false, "Use Japanese wordlist")
	flag.BoolVar(&langCn, "cn", false, "Use Chinese wordlist")
	flag.BoolVar(&langFr, "fr", false, "Use French wordlist")
	flag.BoolVar(&langIt, "it", false, "Use Italian wordlist")
	flag.BoolVar(&langKo, "ko", false, "Use Korean wordlist")
	flag.BoolVar(&langEs, "es", false, "Use Spanish wordlist")

	var useShort bool
	flag.BoolVar(&useShort, "12", false, "Generate 12-word mnemonic (default 24)")

	flag.Parse()

	// Обработка аргументов после флагов
	args := flag.Args()
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-p":
			if i+1 < len(args) {
				password = args[i+1]
				i++ // Пропускаем следующий аргумент
			}
		case "-o":
			if i+1 < len(args) {
				outputFile = args[i+1]
				i++ // Пропускаем следующий аргумент
			}
		case "-s":
			if i+1 < len(args) {
				if size, err := strconv.Atoi(args[i+1]); err == nil {
					qrSize = size
				}
				i++ // Пропускаем следующий аргумент
			}
		}
	}

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

			// Сохраняем в файл если указан
			if outputFile != "" {
				if strings.HasSuffix(outputFile, ".png") {
					if err := qr.GenerateToFile(encrypted, outputFile, qrSize); err != nil {
						log.Printf("Failed to save QR code: %v", err)
					} else {
						fmt.Printf("QR code saved to: %s\n", outputFile)
					}
				}
			}
		} else {
			// Если нет -e, обрабатываем как обычный текст (как с флагом -c)
			result = text
		}
	case useEncrypt:
		// Получаем текст для шифрования
		var text string
		if len(args) > 0 {
			// Собираем все слова до первого флага
			var words []string
			for _, arg := range args {
				if strings.HasPrefix(arg, "-") {
					break
				}
				words = append(words, arg)
			}
			text = strings.Join(words, " ")
		}

		if text == "" {
			log.Fatal("No text provided for encryption. Use: passgen -e \"text\" -p password -o file.png")
		}

		// Проверяем пароль
		if password == "" {
			log.Fatal("Password required for encryption. Use: -p \"password\"")
		}

		// Инициализируем шифрование
		xchacha, err := crypto.NewXChaCha(password)
		if err != nil {
			log.Fatalf("Failed to initialize encryption: %v", err)
		}

		// Шифруем
		encrypted, err := xchacha.Encrypt([]byte(text))
		if err != nil {
			log.Fatalf("Failed to encrypt: %v", err)
		}

		result = encrypted

		// Выводим зашифрованный текст
		fmt.Println(result)

		// Копируем в буфер обмена
		if err := clipboard.Copy(result); err != nil {
			log.Printf("Failed to copy to clipboard: %v", err)
		} else {
			fmt.Println("Copied to clipboard")
		}

		// Если указан файл, сохраняем в него
		if outputFile != "" {
			if err := qr.GenerateToFile(result, outputFile, qrSize); err != nil {
				log.Printf("Failed to save QR code: %v", err)
			} else {
				fmt.Printf("QR code saved to: %s\n", outputFile)
			}
		} else {
			// Если файл не указан, показываем QR в терминале
			qrCode, err := qr.Generate(result)
			if err != nil {
				log.Printf("Failed to create QR code: %v", err)
			} else {
				fmt.Println("\nQR code:")
				fmt.Println(qrCode)
			}
		}
		return // Выходим после обработки шифрования
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
		// Определяем длину (12 или 24 слова)
		wordCount := 24 // По умолчанию
		if useShort {
			wordCount = 12
		}

		// Определяем язык
		lang := "en" // По умолчанию английский
		switch {
		case langRu:
			lang = "ru"
		case langJp:
			lang = "jp"
		case langCn:
			lang = "cn"
		case langFr:
			lang = "fr"
		case langIt:
			lang = "it"
		case langKo:
			lang = "ko"
		case langEs:
			lang = "es"
		}

		bip39Gen, err := bip39.New(lang)
		if err != nil {
			log.Fatalf("Failed to initialize BIP39: %v", err)
		}

		var entropy int
		if wordCount == 12 {
			entropy = bip39.ENT128
		} else {
			entropy = bip39.ENT256
		}

		mnemonic, err := bip39Gen.Generate(entropy)
		if err != nil {
			log.Fatalf("Failed to generate mnemonic: %v", err)
		}
		result = mnemonic
	default:
		// Генерация обычного пароля
		password, err := generator.Generate(length)
		if err != nil {
			log.Fatalf("Failed to generate password: %v", err)
		}
		result = password
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
