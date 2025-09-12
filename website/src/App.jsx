import { useState, useEffect } from "react";
import "./App.css";
import WaveBackground from "./components/WaveBackground";

function App() {
  const [theme, setTheme] = useState("dark");

  const toggleTheme = () => {
    setTheme(theme === "dark" ? "light" : "dark");
  };

  useEffect(() => {
    document.documentElement.setAttribute("data-theme", theme);
  }, [theme]);

  return (
    <div>
      <WaveBackground />
      <div className="container">
        <button className="toggle-btn" onClick={toggleTheme}>
          Switch to {theme === "dark" ? "Light" : "Dark"} Theme
        </button>
        <header>
          <h1>🌐 gochromedp</h1>
          <p>
            A modern command-line tool similar to wkhtmltopdf but powered by
            Chromium headless for superior rendering and web standards support.
          </p>
        </header>

        <section id="features">
          <h2>✨ Features</h2>
          <ul>
            <li>HTML to PDF conversion with full Chromium rendering</li>
            <li>
              HTML to Image conversion (PNG/JPEG) with screenshot capabilities
            </li>
            <li>URL and file input support</li>
            <li>Custom page sizes and margins (A4, A3, Letter, Legal, etc.)</li>
            <li>Page orientation (portrait/landscape)</li>
            <li>High-quality output with modern web standards</li>
            <li>CLI tool similar to wkhtmltopdf interface</li>
            <li>Go library for programmatic use</li>
            <li>Cross-platform (Windows, macOS, Linux)</li>
          </ul>
        </section>

        <section id="installation">
          <h2>🔧 Installation</h2>
          <h3>Prerequisites</h3>
          <ul>
            <li>Go 1.23+</li>
            <li>Chromium/Chrome browser installed</li>
          </ul>
          <h3>Install from source</h3>
          <pre>
            <code>
              git clone https://github.com/chinmay-sawant/gochromedp.git 
              cd gochromedp 
              go mod download go build -o gochromedp ./cmd/gochromedp
            </code>
          </pre>
          <h3>Install pre-built binary</h3>
          <pre>
            <code>
             download from releases page 
              or build and install globally 
              go install github.com/chinmay-sawant/gochromedp/cmd/gochromedp@latest
            </code>
          </pre>
        </section>

        <section id="usage">
          <h2>📖 Usage</h2>
          <h3>Command Line Interface</h3>
          <pre>
            <code>gochromedp [command]</code>
          </pre>
          <p>Available commands: pdf, image, version</p>
          <h3>Convert URL to PDF</h3>
          <pre>
            <code>
              # Basic conversion gochromedp pdf https://example.com output.pdf #
              With custom page size and margins gochromedp pdf --page-size A4
              --margin-top 20mm --margin-bottom 20mm https://example.com
              document.pdf # Landscape orientation gochromedp pdf --orientation
              landscape --page-size A3 https://example.com landscape.pdf
            </code>
          </pre>
          <h3>Convert HTML file to PDF</h3>
          <pre>
            <code>gochromedp pdf input.html output.pdf</code>
          </pre>
          <h3>Convert URL to Image</h3>
          <pre>
            <code>
              # PNG screenshot (default) gochromedp image --width 1920 --height
              1080 https://example.com screenshot.png # JPEG with quality
              setting gochromedp image --format jpeg --quality 85 --width 1024
              --height 768 https://example.com photo.jpg
            </code>
          </pre>
          <h3>Convert HTML file to Image</h3>
          <pre>
            <code>gochromedp image input.html screenshot.png</code>
          </pre>
        </section>

        <section id="options">
          <h2>⚙️ Command Line Options</h2>
          <h3>Global Options</h3>
          <ul>
            <li>
              <code>--page-size string</code> Page size (A4, A3, Letter, Legal)
              (default "A4")
            </li>
            <li>
              <code>--orientation string</code> Page orientation
              (portrait/landscape) (default "portrait")
            </li>
            <li>
              <code>--margin-top string</code> Top margin (default "10mm")
            </li>
            <li>
              <code>--margin-right string</code> Right margin (default "10mm")
            </li>
            <li>
              <code>--margin-bottom string</code> Bottom margin (default "10mm")
            </li>
            <li>
              <code>--margin-left string</code> Left margin (default "10mm")
            </li>
          </ul>
          <h3>PDF Options</h3>
          <ul>
            <li>
              <code>--no-background</code> Do not print background
            </li>
            <li>
              <code>--grayscale</code> Generate grayscale PDF
            </li>
          </ul>
          <h3>Image Options</h3>
          <ul>
            <li>
              <code>--format string</code> Image format (png/jpeg) (default
              "png")
            </li>
            <li>
              <code>--quality int</code> Image quality (1-100, for JPEG)
              (default 90)
            </li>
            <li>
              <code>--width int</code> Viewport width (default 1024)
            </li>
            <li>
              <code>--height int</code> Viewport height (default 768)
            </li>
          </ul>
        </section>

        <section id="library">
          <h2>📚 Go Library Usage</h2>
          <pre>
            <code>{`import (
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
}`}</code>
          </pre>
        </section>

        <section id="examples">
          <h2>🔍 Examples</h2>
          <h3>Basic HTML to PDF</h3>
          <pre>
            <code>gochromedp pdf example.html output.pdf</code>
          </pre>
          <h3>Webpage to high-res screenshot</h3>
          <pre>
            <code>
              gochromedp image --width 1920 --height 1080 --quality 95
              https://github.com screenshot.png
            </code>
          </pre>
          <h3>Custom margins and page size</h3>
          <pre>
            <code>
              gochromedp pdf --page-size Letter --margin-top 25mm
              --margin-bottom 25mm --margin-left 20mm --margin-right 20mm
              document.html print.pdf
            </code>
          </pre>
          <h3>Batch conversion script</h3>
          <pre>
            <code>{`#!/bin/bash
urls=("https://example.com" "https://github.com" "https://golang.org")
for url in "\${urls[@]}"; do
    filename=$(echo $url | sed 's|https://||; s|/|_|g')
    gochromedp pdf "$url" "\${filename}.pdf"
done`}</code>
          </pre>
        </section>

        <section id="architecture">
          <h2>🏗️ Architecture</h2>
          <p>
            gochromedp uses the Chrome DevTools Protocol (CDP) through the{" "}
            <code>chromedp</code> Go library to control a headless Chromium
            instance. This provides:
          </p>
          <ul>
            <li>
              Modern rendering engine with full CSS and JavaScript support
            </li>
            <li>Better font rendering and layout accuracy</li>
            <li>Web standards compliance (ES6+, CSS3, etc.)</li>
            <li>Security through sandboxed browser execution</li>
            <li>Performance optimizations for headless operation</li>
          </ul>
        </section>

        <section id="comparison">
          <h2>🤔 Why gochromedp vs wkhtmltopdf?</h2>
          <table>
            <thead>
              <tr>
                <th>Feature</th>
                <th>wkhtmltopdf</th>
                <th>gochromedp</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Rendering Engine</td>
                <td>Qt WebKit (old)</td>
                <td>Chromium (modern)</td>
              </tr>
              <tr>
                <td>CSS Support</td>
                <td>Limited</td>
                <td>Full CSS3</td>
              </tr>
              <tr>
                <td>JavaScript</td>
                <td>Basic ES5</td>
                <td>Full ES6+</td>
              </tr>
              <tr>
                <td>Fonts</td>
                <td>System fonts only</td>
                <td>Web fonts + system</td>
              </tr>
              <tr>
                <td>Performance</td>
                <td>Fast</td>
                <td>Slightly slower</td>
              </tr>
              <tr>
                <td>Maintenance</td>
                <td>Unmaintained</td>
                <td>Active development</td>
              </tr>
              <tr>
                <td>Dependencies</td>
                <td>Qt libraries</td>
                <td>Chrome/Chromium</td>
              </tr>
            </tbody>
          </table>
        </section>

        <section id="troubleshooting">
          <h2>🐛 Troubleshooting</h2>
          <h3>"Chrome/Chromium not found"</h3>
          <p>Ensure Chrome or Chromium is installed and accessible:</p>
          <pre>
            <code>
              # Linux sudo apt-get install chromium-browser # macOS brew install
              chromium # Windows - Download from https://www.chromium.org/
            </code>
          </pre>
          <h3>"Connection refused" errors</h3>
          <p>
            Try with different Chrome flags or ensure no other Chrome instances
            are running.
          </p>
          <h3>Memory issues</h3>
          <p>For large documents, increase memory limits:</p>
          <pre>
            <code>
              gochromedp pdf --memory-pressure-off large-document.html
              output.pdf
            </code>
          </pre>
        </section>

        <footer>
          <p>MIT License - see LICENSE file for details.</p>
          <p>Built with ❤️ using Go and Chromium</p>
        </footer>
      </div>
    </div>
  );
}

export default App;
