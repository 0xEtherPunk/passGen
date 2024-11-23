package main

import (
	"flag"
	"fmt"
	"log"

	"passgen/internal/bip39"
	"passgen/internal/clipboard"
	"passgen/internal/generator"
	"passgen/internal/qr"
)

func main() {
	var length int
	var useBip39 bool
	var useShort bool
	var lang string
	var outputFile string
	var qrSize int

	flag.IntVar(&length, "length", 0, "Password length (default 24-28)")
	flag.IntVar(&length, "l", 0, "Password length (shorthand)")
	flag.BoolVar(&useBip39, "bip39", false, "Generate BIP39 mnemonic")
	flag.BoolVar(&useBip39, "b", false, "Generate BIP39 mnemonic (shorthand)")
	flag.BoolVar(&useShort, "12", false, "Generate 12-word mnemonic (default is 24)")
	flag.StringVar(&outputFile, "o", "", "Output file for QR code (PNG format)")
	flag.IntVar(&qrSize, "s", 256, "QR code size in pixels (default: 256)")

	// Флаги для языков
	var langRu, langEn, langJp, langFr, langIt, langKo, langCn, langEs bool
	flag.BoolVar(&langRu, "ru", false, "Use Russian wordlist")
	flag.BoolVar(&langEn, "en", false, "Use English wordlist")
	flag.BoolVar(&langJp, "jp", false, "Use Japanese wordlist")
	flag.BoolVar(&langFr, "fr", false, "Use French wordlist")
	flag.BoolVar(&langIt, "it", false, "Use Italian wordlist")
	flag.BoolVar(&langKo, "ko", false, "Use Korean wordlist")
	flag.BoolVar(&langCn, "cn", false, "Use Chinese wordlist")
	flag.BoolVar(&langEs, "es", false, "Use Spanish wordlist")

	flag.Parse()

	var result string
	var err error

	if useBip39 {
		// Определяем язык из флагов
		lang = "en" // по умолчанию
		switch {
		case langRu:
			lang = "ru"
		case langJp:
			lang = "jp"
		case langFr:
			lang = "fr"
		case langIt:
			lang = "it"
		case langKo:
			lang = "ko"
		case langCn:
			lang = "cn"
		case langEs:
			lang = "es"
		}

		mnemo, err := bip39.New(lang)
		if err != nil {
			log.Fatalf("Failed to initialize BIP39: %v", err)
		}

		// Выбираем длину энтропии в зависимости от флага
		entropy := bip39.ENT256 // 24 слова по умолчанию
		if useShort {
			entropy = bip39.ENT128 // 12 слов
		}

		result, err = mnemo.Generate(entropy)
		if err != nil {
			log.Fatalf("Failed to generate mnemonic: %v", err)
		}

		wordCount := "24"
		if useShort {
			wordCount = "12"
		}
		fmt.Printf("Generated BIP39 mnemonic (%s words, %s):\n%s\n", wordCount, lang, result)
	} else {
		gen := generator.New(24, 28)
		result, err = gen.Generate(length)
		if err != nil {
			log.Fatalf("Password generation error: %v", err)
		}
		fmt.Printf("Generated password: %s\n", result)
	}

	if err := clipboard.Copy(result); err != nil {
		log.Printf("Failed to copy to clipboard: %v", err)
	} else {
		fmt.Println("Copied to clipboard")
	}

	if outputFile != "" {
		if err := qr.GenerateToFile(result, outputFile, qrSize); err != nil {
			log.Printf("Failed to save QR code to file: %v", err)
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
