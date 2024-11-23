package qr

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

// Generate creates QR code as ASCII art
func Generate(text string) (string, error) {
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return "", fmt.Errorf("failed to generate QR code: %v", err)
	}
	return qr.ToSmallString(false), nil
}

// GenerateToFile creates QR code and saves it as PNG file
func GenerateToFile(text string, filename string, size int) error {
	if size == 0 {
		size = 256 // default size
	}
	err := qrcode.WriteFile(text, qrcode.Medium, size, filename)
	if err != nil {
		return fmt.Errorf("failed to save QR code to file: %v", err)
	}
	return nil
}
