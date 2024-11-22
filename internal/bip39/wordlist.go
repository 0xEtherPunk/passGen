package bip39

import (
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed wordlist/*.txt
var wordlists embed.FS

// LoadWordlist загружает слова для указанного языка из файла
func LoadWordlist(lang string) ([]string, error) {
	if lang == "" {
		lang = "en"
	}

	data, err := wordlists.ReadFile("wordlist/" + lang + ".txt")
	if err != nil {
		return nil, fmt.Errorf("failed to load wordlist: %v", err)
	}

	words := strings.Split(string(data), "\n")
	var cleanWords []string
	for _, word := range words {
		if w := strings.TrimSpace(word); w != "" {
			cleanWords = append(cleanWords, w)
		}
	}

	return cleanWords, nil
}

// ListLanguages возвращает список доступных языков
func ListLanguages() ([]string, error) {
	files, err := os.ReadDir("wordlist")
	if err != nil {
		return nil, fmt.Errorf("failed to read wordlist directory: %v", err)
	}

	var languages []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			lang := strings.TrimSuffix(file.Name(), ".txt")
			languages = append(languages, lang)
		}
	}

	return languages, nil
}
