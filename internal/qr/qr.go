package qr

import (
	"bytes"

	"github.com/skip2/go-qrcode"
)

// Generate creates QR code from string and returns it as terminal-friendly string
func Generate(text string) (string, error) {
	qr, err := qrcode.New(text, qrcode.Low)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	grid := qr.Bitmap()

	// Находим границы непустой области QR-кода
	minX, maxX, minY, maxY := findBounds(grid)

	// Вычисляем ширину содержимого (каждый пиксель = 2 символа)
	width := (maxX - minX) * 2

	// Верхняя граница с правильной шириной
	buf.WriteString("┌")
	for i := 0; i < width; i++ {
		buf.WriteString("─")
	}
	buf.WriteString("┐\n")

	// QR код
	for y := minY; y < maxY; y++ {
		buf.WriteString("│")
		for x := minX; x < maxX; x++ {
			if grid[y][x] {
				buf.WriteString("██")
			} else {
				buf.WriteString("  ")
			}
		}
		buf.WriteString("│\n")
	}

	// Нижняя граница с правильной шириной
	buf.WriteString("└")
	for i := 0; i < width; i++ {
		buf.WriteString("─")
	}
	buf.WriteString("┘\n")

	return buf.String(), nil
}

// findBounds находит реальные границы QR-кода без пустых областей
func findBounds(grid [][]bool) (minX, maxX, minY, maxY int) {
	minX, minY = len(grid), len(grid)
	maxX, maxY = 0, 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// Добавляем небольшой отступ в 1 символ
	if minX > 0 {
		minX--
	}
	if maxX < len(grid)-1 {
		maxX++
	}
	if minY > 0 {
		minY--
	}
	if maxY < len(grid)-1 {
		maxY++
	}

	return minX, maxX + 1, minY, maxY + 1
}
