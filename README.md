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

#### 🌐 Supported Languages
```bash
# Full phrases (24 words)
passgen -b -en     # 🇬🇧 English (default)
passgen -b -ru     # 🇷🇺 Russian (Русский)
passgen -b -jp     # 🇯🇵 Japanese (日本語)
passgen -b -cn     # 🇨🇳 Chinese (简体中文)
passgen -b -fr     # 🇫🇷 French (Français)
passgen -b -it     # 🇮🇹 Italian (Italiano)
passgen -b -ko     # 🇰🇷 Korean (한국어)
passgen -b -es     # 🇪🇸 Spanish (Español)

# Short phrases (12 words)
passgen -b -12 -en    # 🇬🇧 English
passgen -b -12 -ru    # 🇷🇺 Русский
passgen -b -12 -jp    # 🇯🇵 日本語
passgen -b -12 -cn    # 🇨🇳 简体中文
passgen -b -12 -fr    # 🇫🇷 Français
passgen -b -12 -it    # 🇮🇹 Italiano
passgen -b -12 -ko    # 🇰🇷 한국어
passgen -b -12 -es    # 🇪🇸 Español
```

### 📤 Output Features
Every generated password or mnemonic is automatically:
- 📝 Displayed in terminal
- 📋 Copied to clipboard
- 📱 Converted to QR code

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