<div align="center">

# 🔐 PassGen

### Cryptographically Secure Password Generator with QR Codes

[![Go Version](https://img.shields.io/badge/Go-1.23.2-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

</div>

## 🌟 About

PassGen isn't just another password generator. It's a powerful command-line tool that combines Unix system security with modern features for convenient use. Using a cryptographically secure random number generator (/dev/urandom), PassGen creates truly random passwords and makes them easily accessible through clipboard and QR codes.

## ✨ Features

- 🔒 **Maximum Security**: Utilizing /dev/urandom for cryptographic strength
- 📏 **Flexible Length**: Default 24-28 characters with customization options
- 📋 **Smart Clipboard**: Instant copying without extra characters
- 📱 **Terminal QR Codes**: Quick transfer to mobile devices
- 🎯 **Simple Interface**: Minimalistic and intuitive CLI
- 💻 **Cross-Platform**: Works on Linux, macOS, and Windows

## 🚀 Installation

```bash
go install github.com/0xEtherPunk/passGen/cmd/passgen@latest
```

## 🎮 Usage

### Basic Usage
```bash
passgen
```

### With Custom Length
```bash
passgen -l 32        # short version
passgen -length 32   # full version
```

## 🔧 How It Works

PassGen uses a powerful Unix pipeline for password generation:

```bash
strings /dev/urandom | tr -d "\n[:space:]" | fold -w$((24 + $(od -An -N2 -i /dev/urandom) % 4)) | head -n1
```

This process:
1. 📖 Extracts readable strings from /dev/urandom
2. 🧹 Cleans spaces and newlines
3. 📏 Formats to desired length
4. ✂️ Selects the first line of output

## 🏗 Project Structure

```
passGen/
├── cmd/
│   └── passgen/
│       └── main.go       # Entry point
├── internal/
│   ├── generator/        # Password generation
│   ├── clipboard/        # Clipboard operations
│   └── qr/              # QR code generator
└── README.md
```

## 📦 Dependencies

- [atotto/clipboard](https://github.com/atotto/clipboard) - Cross-platform clipboard
- [skip2/go-qrcode](https://github.com/skip2/go-qrcode) - QR code generation

## 🛠 Requirements

- Go 1.23.2+
- Unix-like system (for /dev/urandom)
- UTF-8 terminal support

## 🤝 Contributing

Your contributions are welcome! Here's how you can help:

1. 🍴 Fork the repository
2. 🌿 Create your feature branch (`git checkout -b feature/amazing`)
3. 🔧 Make your changes
4. 📝 Submit a Pull Request

## 📜 License

MIT © [0xEtherPunk](https://github.com/0xEtherPunk)

## 💖 Acknowledgments

Special thanks to the Unix community for the inspiration and tools that make this project possible.

---

<div align="center">
  
### Made with ❤️ for the Community

</div>