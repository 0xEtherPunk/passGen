package qr

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"strings"

	"github.com/makiuchi-d/gozxing"
	gozxingqr "github.com/makiuchi-d/gozxing/qrcode"
	goqrcode "github.com/skip2/go-qrcode"
)

// Generate creates QR code as ASCII art
func Generate(text string) (string, error) {
	qr, err := goqrcode.New(text, goqrcode.Low)
	if err != nil {
		return "", fmt.Errorf("failed to generate QR code: %v", err)
	}
	return qr.ToSmallString(false), nil
}

// GenerateToFile creates QR code and saves it as PNG file
func GenerateToFile(text string, filename string, size int) error {
	if size == 0 {
		size = 256
	}
	err := goqrcode.WriteFile(text, goqrcode.Low, size, filename)
	if err != nil {
		return fmt.Errorf("failed to save QR code to file: %v", err)
	}
	return nil
}

// ReadFromFile читает данные из файла, автоматически определяя формат
func ReadFromFile(filename string) (string, error) {
	if strings.HasSuffix(filename, ".png") {
		return readQRFromPNG(filename)
	}
	// Для других форматов читаем как текст
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}
	return string(data), nil
}

// readQRFromPNG читает QR код из PNG файла
func readQRFromPNG(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %v", err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", fmt.Errorf("failed to create bitmap: %v", err)
	}

	qrReader := gozxingqr.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", fmt.Errorf("failed to read QR code: %v", err)
	}

	return result.String(), nil
}
