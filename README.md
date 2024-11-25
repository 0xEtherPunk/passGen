<div align="center">

# 🔐 PassGen

### Secure Password & BIP39 Mnemonic Generator

[![Go Version](https://img.shields.io/badge/Go-1.23.2-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

![Demo](demo.gif)

</div>

## 🎯 Overview

PassGen combines secure password generation with BIP39 mnemonic phrase support and encryption capabilities:
- 🛡️ Cryptographically secure generation using /dev/urandom
- 🌍 Multi-language BIP39 support
- 🔒 XChaCha20-Poly1305 encryption
- 📱 QR code generation and reading
- 📋 Instant clipboard integration

## ⚡ Quick Start

```bash
go install github.com/0xEtherPunk/passGen/cmd/passgen@latest
```

## 🛠️ Usage

### 🔑 Password Generation
```bash
# Generate password (24-28 characters)
passgen

# Custom length password
passgen -l 32
```

### 🎲 BIP39 Mnemonic Generation
```bash
# English (default, 24 words)
passgen -b

# Short version (12 words)
passgen -b -12

# Available languages:
passgen -b -en     # 🇬🇧 English (default)
passgen -b -ru     # 🇷🇺 Russian (Русский)
passgen -b -jp     # 🇯🇵 Japanese (日本語)
passgen -b -cn     # 🇨🇳 Chinese (简体中文)
passgen -b -fr     # 🇫🇷 French (Français)
passgen -b -it     # 🇮🇹 Italian (Italiano)
passgen -b -ko     # 🇰🇷 Korean (한국어)
passgen -b -es     # 🇪🇸 Spanish (Español)
```

### 🔐 Encryption & QR Codes
```bash
# Encrypt text and generate QR code
passgen -e "secret message" -p "mypassword" -o secret.png

# Decrypt from QR code
passgen -d secret.png -p "mypassword"

# Encrypt with custom QR size (default: 256x256)
passgen -e "secret" -p "pass" -o large.png -s 512

# Pipe support
echo "secret text" | passgen -e -p "pass" -o qr.png
cat file.txt | passgen -e -p "pass" -o qr.png
```

### 🔍 Help Command
```bash
passgen -help
passgen -h
```

## 📤 Output Features
Every generated output is automatically:
- 📝 Displayed in terminal
- 📋 Copied to clipboard
- 📱 Generated as QR code (if -o flag is used)

## 🏗️ Project Structure
```
passGen/
├── cmd/
│   └── passgen/
│       └── main.go           # 🎯 Entry point
├── internal/
│   ├── bip39/               # 🎲 BIP39 implementation
│   │   ├── wordlist/        # 🌐 Language wordlists
│   │   ├── bip39.go        
│   │   └── wordlist.go     
│   ├── crypto/              # 🔒 Encryption
│   │   └── xchacha.go       # XChaCha20-Poly1305
│   ├── clipboard/           # 📋 Clipboard operations
│   ├── generator/           # 🎯 Password generation
│   └── qr/                  # 📱 QR code operations
└── README.md
```

## ⚙️ Requirements
- 🔧 Go 1.23.2 or higher
- 🐧 Unix-like system (for /dev/urandom)

## 📄 License
MIT © [0xEtherPunk](https://github.com/0xEtherPunk)

---

<div align="center">

### 🌟 If you find PassGen useful, please star it on GitHub!

[![GitHub stars](https://img.shields.io/github/stars/0xEtherPunk/passGen?style=social)](https://github.com/0xEtherPunk/passGen)

</div>

> 🌈 **Pro tip**: Pipe the output through `lolcat` for some extra color magic:
> ```bash
> passgen -b | lolcat
> passgen -e "secret" -p "pass" | lolcat
> ```