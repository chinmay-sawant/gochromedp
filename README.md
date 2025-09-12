# 🌐 gochromedp

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

> 🚀 A modern command-line tool similar to wkhtmltopdf but powered by Chromium headless for superior rendering and web standards support.

## ✨ Features

- **HTML to PDF conversion** with full Chromium rendering
- **HTML to Image conversion** (PNG/JPEG) with screenshot capabilities
- **URL and file input support**
- **Custom page sizes and margins** (A4, A3, Letter, Legal, etc.)
- **Page orientation** (portrait/landscape)
- **High-quality output** with modern web standards
- **CLI tool** similar to wkhtmltopdf interface
- **Go library** for programmatic use
- **Cross-platform** (Windows, macOS, Linux)

## 🔧 Installation

### Prerequisites
- Go 1.23+
- Chromium/Chrome browser installed

### Install from source
```bash
git clone https://github.com/chinmay-sawant/gochromedp.git
cd gochromedp
go mod download
go build -o gochromedp ./cmd/gochromedp
```

### Install pre-built binary
```bash
# Download from releases page
# Or build and install globally
go install github.com/chinmay-sawant/gochromedp/cmd/gochromedp@latest
```

## 📖 Usage

### Command Line Interface

```bash
gochromedp [command]
```

Available commands:
- `pdf` - Convert HTML to PDF
- `image` - Convert HTML to image
- `version` - Print version information

Use `gochromedp [command] --help` for more information about a command.

### Convert URL to PDF
```bash
# Basic conversion
gochromedp pdf https://example.com output.pdf

# With custom page size and margins
gochromedp pdf --page-size A4 --margin-top 20mm --margin-bottom 20mm https://example.com document.pdf

# Landscape orientation
gochromedp pdf --orientation landscape --page-size A3 https://example.com landscape.pdf
```

### Convert HTML file to PDF
```bash
gochromedp pdf input.html output.pdf
```

### Convert URL to Image
```bash
# PNG screenshot (default)
gochromedp image --width 1920 --height 1080 https://example.com screenshot.png

# JPEG with quality setting
gochromedp image --format jpeg --quality 85 --width 1024 --height 768 https://example.com photo.jpg
```

### Convert HTML file to Image
```bash
gochromedp image input.html screenshot.png
```

## ⚙️ Command Line Options

### Global Options
- `--page-size string`     Page size (A4, A3, Letter, Legal) (default "A4")
- `--orientation string`   Page orientation (portrait/landscape) (default "portrait")
- `--margin-top string`    Top margin (default "10mm")
- `--margin-right string`  Right margin (default "10mm")
- `--margin-bottom string` Bottom margin (default "10mm")
- `--margin-left string`   Left margin (default "10mm")

### PDF Options
- `--no-background`        Do not print background
- `--grayscale`           Generate grayscale PDF

### Image Options
- `--format string`       Image format (png/jpeg) (default "png")
- `--quality int`         Image quality (1-100, for JPEG) (default 90)
- `--width int`           Viewport width (default 1024)
- `--height int`          Viewport height (default 768)

## 📚 Go Library Usage

```go
package main

import (
    "os"
    "github.com/chinmay-sawant/gochromedp"
)

func main() {
    // Convert HTML to PDF
    html := "<html><body><h1>Hello World!</h1></body></html>"
    options := &gochromedp.ConvertOptions{
        PageSize:    "A4",
        Orientation: "portrait",
        MarginTop:   "10mm",
    }

    pdfData, err := gochromedp.ConvertHTMLToPDF(html, options)
    if err != nil {
        panic(err)
    }

    os.WriteFile("output.pdf", pdfData, 0644)

    // Convert URL to PDF
    pdfData, err = gochromedp.ConvertURLToPDF("https://example.com", options)
    if err != nil {
        panic(err)
    }

    os.WriteFile("webpage.pdf", pdfData, 0644)

    // Convert HTML to image
    imageData, err := gochromedp.ConvertHTMLToImage(html, &gochromedp.ConvertOptions{
        Format: "png",
        Width:  1024,
        Height: 768,
    })
    if err != nil {
        panic(err)
    }

    os.WriteFile("html-screenshot.png", imageData, 0644)

    // Convert URL to image
    imageData, err = gochromedp.ConvertURLToImage("https://example.com", &gochromedp.ConvertOptions{
        Format: "png",
        Width:  1024,
        Height: 768,
    })
    if err != nil {
        panic(err)
    }

    os.WriteFile("screenshot.png", imageData, 0644)
}
```

## 🔍 Examples

### Basic HTML to PDF
```bash
gochromedp pdf example.html output.pdf
```

### Webpage to high-res screenshot
```bash
gochromedp image --width 1920 --height 1080 --quality 95 https://github.com screenshot.png
```

### Custom margins and page size
```bash
gochromedp pdf --page-size Letter --margin-top 25mm --margin-bottom 25mm --margin-left 20mm --margin-right 20mm document.html print.pdf
```

### Batch conversion script
```bash
#!/bin/bash
urls=("https://example.com" "https://github.com" "https://golang.org")
for url in "${urls[@]}"; do
    filename=$(echo $url | sed 's|https://||; s|/|_|g')
    gochromedp pdf "$url" "${filename}.pdf"
done
```

## 🏗️ Architecture

gochromedp uses the Chrome DevTools Protocol (CDP) through the `chromedp` Go library to control a headless Chromium instance. This provides:

- **Modern rendering engine** with full CSS and JavaScript support
- **Better font rendering** and layout accuracy
- **Web standards compliance** (ES6+, CSS3, etc.)
- **Security** through sandboxed browser execution
- **Performance** optimizations for headless operation

## 🤔 Why gochromedp vs wkhtmltopdf?

| Feature | wkhtmltopdf | gochromedp |
|---------|-------------|-------------|
| Rendering Engine | Qt WebKit (old) | Chromium (modern) |
| CSS Support | Limited | Full CSS3 |
| JavaScript | Basic ES5 | Full ES6+ |
| Fonts | System fonts only | Web fonts + system |
| Performance | Fast | Slightly slower |
| Maintenance | Unmaintained | Active development |
| Dependencies | Qt libraries | Chrome/Chromium |

## 🐛 Troubleshooting

### "Chrome/Chromium not found"
Ensure Chrome or Chromium is installed and accessible:
```bash
# Linux
sudo apt-get install chromium-browser

# macOS
brew install chromium

# Windows - Download from https://www.chromium.org/
```

### "Connection refused" errors
Try with different Chrome flags or ensure no other Chrome instances are running.

### Memory issues
For large documents, increase memory limits:
```bash
gochromedp pdf --memory-pressure-off large-document.html output.pdf
```

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 🙏 Acknowledgments

- [chromedp](https://github.com/chromedp/chromedp) - Chrome DevTools Protocol library
- [cobra](https://github.com/spf13/cobra) - CLI framework
- Inspired by [wkhtmltopdf](https://wkhtmltopdf.org/)