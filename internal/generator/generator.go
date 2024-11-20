package generator

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

type Generator struct {
	minLength int
	maxLength int
}

func New(minLength, maxLength int) *Generator {
	return &Generator{
		minLength: minLength,
		maxLength: maxLength,
	}
}

// Generate creates a new password of specified length
func (g *Generator) Generate(length int) (string, error) {
	var cmd *exec.Cmd

	if length == 0 {
		// Using original algorithm with random length
		cmd = exec.Command("bash", "-c",
			`strings /dev/urandom | tr -d "\n[:space:]" | fold -w$((24 + $(od -An -N2 -i /dev/urandom) % 4)) | head -n1 | tr -d "\n"`)
	} else {
		// Using specified length
		cmd = exec.Command("bash", "-c",
			"strings /dev/urandom | tr -d '\n[:space:]' | fold -w"+strconv.Itoa(length)+" | head -n1 | tr -d '\n'")
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(out.String(), "\n"), nil
}
