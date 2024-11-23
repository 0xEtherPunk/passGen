<div align="center">

# 🔐 PassGen

### Secure Password & BIP39 Mnemonic Generator

[![Go Version](https://img.shields.io/badge/Go-1.23.2-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

![Demo](demo.gif)

</div>

## 🎯 Overview

PassGen combines secure password generation with BIP39 mnemonic phrase support, offering:
- 🛡️ Cryptographically secure generation using /dev/urandom
- 🌍 Multi-language BIP39 support
- 📋 Instant clipboard integration
- 📱 Terminal QR code display

## ⚡ Quick Start

### Installation Options

1. Latest release (recommended):
```bash
go install github.com/0xEtherPunk/passGen/cmd/passgen@latest
```

2. Specific version:
```bash
go install github.com/0xEtherPunk/passGen/cmd/passgen@v1.1.0
```

3. From source:
```bash
# Clone repository
git clone https://github.com/0xEtherPunk/passGen.git
cd passGen

# Install locally
go install ./cmd/passgen
```

4. Manual build:
```bash
git clone https://github.com/0xEtherPunk/passGen.git
cd passGen
go build -o passgen cmd/passgen/main.go
sudo mv passgen /usr/local/bin/  # Optional: make globally available
```

## 🛠️ Usage

### 🔑 Password Generation
```bash
# Standard password (24-28 characters)
passgen

# Custom length password
passgen -l 32
```

### 🎲 BIP39 Mnemonic Generation

#### Default Usage
```bash
passgen -b         # 24 words in English
passgen -b -12     # 12 words in English
```

### 🔍 Help Command
```bash
# Show all available options
passgen -help
passgen -h

# Output includes:
  -12
        Generate 12-word mnemonic (default is 24)
  -b, -bip39
        Generate BIP39 mnemonic
  -l, -length int
        Password length (default 24-28)
  -o string
        Output file for QR code (PNG format)
  -s int
        QR code size in pixels (default: 256)
# Language options:
  -en    Use English wordlist (default)    # 🇬🇧
  -ru    Use Russian wordlist              # 🇷🇺
  -jp    Use Japanese wordlist             # 🇯🇵
  -cn    Use Chinese wordlist              # 🇨🇳
  -fr    Use French wordlist               # 🇫🇷
  -it    Use Italian wordlist              # 🇮🇹
  -ko    Use Korean wordlist               # 🇰🇷
  -es    Use Spanish wordlist              # 🇪🇸
```

### 📤 Output Features
Every generated password or mnemonic is automatically:
- 📝 Displayed in terminal
- 📋 Copied to clipboard
- 📱 Converted to QR code

### QR Code Options
```bash
# Display QR in terminal (default)
passgen -b

# Save QR as PNG file (default size: 256x256)
passgen -b -o mnemonic.png

# Save QR with custom size (in pixels)
passgen -b -o mnemonic.png -s 512    # 512x512
passgen -b -o mnemonic.png -s 1024   # 1024x1024

# Save QR with custom path
passgen -b -o ~/Documents/mnemonic.png
passgen -b -o ../backup/phrase.png

# Examples with different options combined
passgen -b -12 -ru -o russian-12words.png -s 512     # Russian 12-word phrase, 512x512 QR
passgen -b -jp -o ~/backup/japanese.png -s 1024      # Japanese 24-word phrase, 1024x1024 QR
passgen -l 32 -o password.png                        # 32-char password QR
```

All outputs (password/mnemonic) are still:
- 📝 Displayed in terminal
- 📋 Copied to clipboard
- 💾 Saved as QR code (if -o flag is used)

## 🏗️ Project Structure
```
passGen/
├── cmd/
│   └── passgen/
│       └── main.go           # 🎯 Entry point
├── internal/
│   ├── bip39/               # 🎲 BIP39 implementation
│   │   ├── wordlist/        # 🌐 Language wordlists
│   │   │   ├── en.txt      # English
│   │   │   ├── ru.txt      # Russian
│   │   │   ├── jp.txt      # Japanese
│   │   │   ├── cn.txt      # Chinese
│   │   │   ├── fr.txt      # French
│   │   │   ├── it.txt      # Italian
│   │   │   ├── ko.txt      # Korean
│   │   │   └── es.txt      # Spanish
│   │   ├── bip39.go        # Core BIP39 logic
│   │   └── wordlist.go     # Wordlist handling
│   ├── clipboard/           # 📋 Clipboard operations
│   ├── generator/           # 🎯 Password generation
│   └── qr/                  # 📱 QR code generation
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

</div>

---

> 🌈 **Pro tip**: Pipe the output through `lolcat` for some extra color magic:
> ```bash
> passgen | lolcat
> passgen -b -12 -cn | lolcat
> ```