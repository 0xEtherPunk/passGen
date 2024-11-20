package main

import (
	"flag"
	"fmt"
	"log"

	"passgen/internal/clipboard"
	"passgen/internal/generator"
	"passgen/internal/qr"
)

func main() {
	var length int
	flag.IntVar(&length, "length", 0, "Password length (default 24-28)")
	flag.IntVar(&length, "l", 0, "Password length (shorthand)")
	flag.Parse()

	gen := generator.New(24, 28)

	password, err := gen.Generate(length)
	if err != nil {
		log.Fatalf("Password generation error: %v", err)
	}

	fmt.Printf("Generated password: %s\n", password)

	if err := clipboard.Copy(password); err != nil {
		log.Printf("Failed to copy password to clipboard: %v", err)
	} else {
		fmt.Println("Password copied to clipboard")
	}

	qrCode, err := qr.Generate(password)
	if err != nil {
		log.Printf("Failed to create QR code: %v", err)
	} else {
		fmt.Println("\nQR code for password:")
		fmt.Println(qrCode)
	}
}
