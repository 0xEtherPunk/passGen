#!/bin/bash

# Create main directory structure
mkdir -p cmd/passgen
mkdir -p internal/{bip39,clipboard,generator,qr}
mkdir -p internal/bip39/wordlist

# Create main files
touch cmd/passgen/main.go
touch internal/clipboard/clipboard.go
touch internal/generator/generator.go
touch internal/qr/qr.go
touch internal/bip39/{bip39.go,wordlist.go}

# Create wordlist files
touch internal/bip39/wordlist/{en,ru,jp,cn,fr,it,ko,es}.txt

# Create root files
touch go.mod
touch README.md
touch LICENSE 