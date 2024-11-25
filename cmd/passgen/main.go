package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"passgen/internal/clipboard"
	"passgen/internal/qr"
)

func main() {
	var length int
	var useBip39 bool
	var customText string
	var outputFile string
	var qrSize int
	var usePipe bool

	flag.IntVar(&length, "l", 0, "Password length (shorthand)")
	flag.BoolVar(&useBip39, "b", false, "Generate BIP39 mnemonic")
	flag.StringVar(&customText, "c", "", "Use custom text for QR code")
	flag.StringVar(&outputFile, "o", "", "Output file for QR code (PNG format)")
	flag.IntVar(&qrSize, "s", 256, "QR code size in pixels (default: 256)")
	flag.Parse()

	// Проверяем pipe ввод
	stat, _ := os.Stdin.Stat()
	usePipe = (stat.Mode() & os.ModeCharDevice) == 0

	var result string

	switch {
	case usePipe:
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Failed to read from pipe: %v", err)
		}
		result = strings.TrimSpace(string(bytes))
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
