package generator

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Generate creates a new password using original algorithm
func Generate(length int) (string, error) {
	var cmd *exec.Cmd

	if length == 0 {
		// Оригинальный алгоритм с случайной длиной
		cmd = exec.Command("bash", "-c",
			`strings /dev/urandom | tr -d "\n[:space:]" | fold -w$((24 + $(od -An -N2 -i /dev/urandom) % 4)) | head -n1`)
	} else {
		// Фиксированная длина
		cmd = exec.Command("bash", "-c",
			`strings /dev/urandom | tr -d "\n[:space:]" | fold -w`+fmt.Sprintf("%d", length)+` | head -n1`)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
