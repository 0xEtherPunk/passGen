package clipboard

import "github.com/atotto/clipboard"

// Copy copies text to clipboard
func Copy(text string) error {
	return clipboard.WriteAll(text)
}
