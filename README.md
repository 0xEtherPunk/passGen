<div align="center">

# 🔐 PassGen

### Secure Password Generator & Encryption Tool

[![Go Version](https://img.shields.io/badge/Go-1.23.2-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

![Demo](demo.gif)

</div>

## 🚀 Features
- 🎲 Cryptographically secure password generation
- 🔒 XChaCha20-Poly1305 encryption
- 🌍 Multi-language BIP39 mnemonic support
- 📱 QR code generation and reading
- 📋 Automatic clipboard integration
- 📤 Pipe support for text input/output

## 📦 Installation
```bash
go install github.com/username/passgen@latest

# Optional: Create alias
echo 'alias pg="passgen"' >> ~/.bashrc
```

## 🎯 Command Flags
### Basic Flags
- `-l <number>` - Set password length (default: random 24-28)
- `-o <file>` - Save output to PNG file
- `-s <size>` - Set QR code size in pixels (default: 256)

### Encryption Flags
- `-e <text>` - Encrypt text (requires -p)
- `-p <password>` - Password for encryption/decryption
- `-d <file/text>` - Decrypt from file or text

### BIP39 Flags
- `-b` - Generate BIP39 mnemonic (24 words by default)
- `-12` - Generate 12-word mnemonic (use with -b)

### Language Flags (for BIP39)
- `-en` - English wordlist (default)
- `-ru` - Russian wordlist 🇷🇺
- `-jp` - Japanese wordlist 🇯🇵
- `-cn` - Chinese wordlist 🇨🇳
- `-fr` - French wordlist 🇫🇷
- `-it` - Italian wordlist 🇮🇹
- `-ko` - Korean wordlist 🇰🇷
- `-es` - Spanish wordlist 🇪🇸

### Custom Flag
- `-c <text>` - Create QR code from custom text

### Examples
```bash
# Password generation
pg -l 32                    # 32-char password
pg -l 16 -o pass.png       # 16-char password with QR

# Encryption
pg -e secret -p pass   # Encrypt text
pg -d file.png -p pass   # Decrypt from file

# BIP39
pg -b                      # 24 words in English
pg -b -12 -ru             # 12 words in Russian
pg -b -jp -o seed.png     # Japanese with QR
```

## 🛠️ Usage Examples

### 🎲 Password Generation
```bash
# Basic password (24-28 chars)
pg
pg -o pass.png           # Save as QR
pg -s 512 -o pass.png   # Custom QR size

# Custom length
pg -l 32
pg -l 16 -o pass.png
```

### 🔐 Encryption

![Encryption](crypto.gif)

```bash
# Basic encryption
pg -e secret text -p password123
pg -e secret text -p password123 -o secret.png

# Multi-word text
pg -e this is my secret text -p pass123 -o secret.png

# Using generated password from clipboard
pg -o pass.png                   # Generate and save password
pg -e secret text -p "$(xclip)" -o secret.png

# Pipe input
echo "secret text" | pg -e -p "pass123"
cat file.txt | pg -e -p "pass123" -o encrypted.png

# Custom QR sizes
pg -e "secret" -p "pass" -o large.png -s 512
pg -e "secret" -p "pass" -o huge.png -s 1024
```

### 🔓 Decryption
```bash
# From QR file
pg -d secret.png -p "pass123"

# From encrypted text
pg -d "encrypted_base64_text" -p "pass123"

# Save decrypted to file
pg -d secret.png -p "pass123" > decrypted.txt
```

### 🌍 BIP39 Mnemonic Generation
```bash
# English (default)
pg -b            # 24 words
pg -b -12        # 12 words
pg -b -o mnemonic.png

# Other languages
pg -b -ru        # 🇷🇺 Russian
pg -b -jp        # 🇯🇵 Japanese
pg -b -cn        # 🇨🇳 Chinese
pg -b -fr        # 🇫🇷 French
pg -b -it        # 🇮🇹 Italian
pg -b -ko        # 🇰🇷 Korean
pg -b -es        # 🇪🇸 Spanish

# Combined flags
pg -b -12 -ru -o mnemonic.png    # 12 Russian words with QR
pg -b -jp -s 512 -o phrase.png   # Japanese with large QR
```

### 📱 QR Code Operations
```bash
# Custom text to QR
pg -c "any text" -o qr.png
pg -c "large text" -s 512 -o qr.png

# Read from QR
pg -d qr.png
```

### 🔄 Pipeline Examples
```bash
# Generate BIP39 and encrypt with clipboard password
pg -l 32                    # Generate and copy password
pg -b -12 -o seed.png | pg -e -p "$(xclip -o)" -o backup.png

# Or using xsel
pg -b -12 -o seed.png | pg -e -p "$(xsel -b)" -o backup.png

# For macOS:
pg -b -12 -o seed.png | pg -e -p "$(pbpaste)" -o backup.png

### 🔄 Advanced Usage
```bash
# Encrypt BIP39 phrase
pg -b -12 -o seed.png | pg -e -p "pass123" -o backup.png

# Create encrypted backup
tar czf - documents/ | \
  pg -e -p "pass123" -o backup.png -s 1024
```

### 🎨 Creative Use Cases
```bash
# Secure BIP39 backup with encryption
pg -b -12 -o seed.png | pg -e -p "secret123" -o encrypted_seed.png -s 1000

# Multi-language secure backup
pg -b -12 -ru -o seed_ru.png | pg -e -p "пароль123" -o backup_ru.png
pg -b -12 -jp -o seed_jp.png | pg -e -p "パスワード" -o backup_jp.png

# Create encrypted archive with seeds
mkdir seeds/
pg -b -12 -o seeds/en.png
pg -b -12 -ru -o seeds/ru.png
pg -b -12 -jp -o seeds/jp.png
tar czf - seeds/ | pg -e -p "archive123" -o seeds_backup.png -s 2000

# Secure password sharing
pg -l 32 -o pass.png | pg -e -p "share123" -o shared_pass.png
# Recipient can decrypt with: pg -d shared_pass.png -p "share123"
```

## 🔍 Tips & Tricks
- 🎯 Generated passwords are automatically copied to clipboard
- 🖼️ QR codes are shown in terminal if no output file specified
- 📋 Encrypted text is copied to clipboard for easy sharing
- 🔄 Pipe support works with any text-producing command
- 🎨 Custom QR sizes help with scanning distance/resolution

### 💡 Tips & Tricks
- Generate and encrypt in one command using pipes
- Use different QR sizes for different data lengths
- Combine BIP39 languages for extra entropy
- Store encryption keys as separate QR codes
- Use generated passwords for encryption
- Create multi-part backups for extra security

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
- 📋 xclip/xsel for Linux clipboard support
- 📋 pbcopy/pbpaste for macOS clipboard support

## 🔒 Technical Details
### Password Generation
- Uses /dev/urandom for cryptographic randomness
- Default length: 24-28 characters
- Character set includes:
  - Lowercase letters (a-z)
  - Uppercase letters (A-Z)
  - Numbers (0-9)
  - Special characters (!@#$%^&*()_+-=[]{}|;:,.<>?)

### BIP39 Implementation
- Supports 8 languages: 🇺🇸 EN, 🇷🇺 RU, 🇯🇵 JP, 🇨🇳 CN, 🇫🇷 FR, 🇮🇹 IT, 🇰🇷 KO, 🇪🇸 ES
- 12 or 24 word phrases
- Follows official BIP39 specification
- Entropy: 128 bits (12 words) or 256 bits (24 words)

### Encryption Details
- Algorithm: XChaCha20-Poly1305
- Unique salt for each encryption

### QR Code Features
- Default size: 256x256 pixels
- Custom sizes supported
- Supports both generation and reading
- ASCII art display in terminal

---

<div align="center">

### 🌟 If you find PassGen useful, please star it on GitHub!

</div>