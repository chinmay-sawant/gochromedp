# gochromedp Makefile

.PHONY: all build clean test install deps examples

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=gochromedp
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the project
all: deps build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/gochromedp

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/gochromedp

# Build for Windows
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME).exe -v ./cmd/gochromedp

# Build for macOS
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_darwin -v ./cmd/gochromedp

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_NAME).exe
	rm -f $(BINARY_NAME)_darwin

# Run tests
test:
	$(GOTEST) -v ./...

# Download dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Install the binary globally
install: build
	sudo cp $(BINARY_NAME) /usr/local/bin/

# Generate examples
examples: build
	@echo "Generating PDF example..."
	./$(BINARY_NAME) pdf examples/demo.html examples/demo.pdf
	@echo "Generating image example..."
	./$(BINARY_NAME) image --width 1200 --height 800 examples/demo.html examples/demo.png
	@echo "Examples generated in examples/ directory"

# Run example conversions
demo: build
	@echo "Converting demo HTML to PDF..."
	./$(BINARY_NAME) pdf --page-size A4 --margin-top 20mm examples/demo.html demo.pdf
	@echo "Converting demo HTML to PNG..."
	./$(BINARY_NAME) image --width 1024 --height 768 examples/demo.html demo.png
	@echo "Demo files generated: demo.pdf, demo.png"

# Development setup
dev-setup: deps
	@echo "Development environment ready!"

# Cross-platform build
cross-build: build-linux build-windows build-darwin
	@echo "Cross-platform builds completed"

# Package for distribution
package: build
	mkdir -p dist
	cp $(BINARY_NAME) dist/
	cp README.md dist/
	cp LICENSE dist/
	tar -czf dist/$(BINARY_NAME)-$(shell git describe --tags --abbrev=0 2>/dev/null || echo "v1.0.0").tar.gz -C dist .

# Help
help:
	@echo "Available targets:"
	@echo "  all          - Download deps and build"
	@echo "  build        - Build the binary"
	@echo "  build-linux  - Build for Linux"
	@echo "  build-windows- Build for Windows"
	@echo "  build-darwin - Build for macOS"
	@echo "  clean        - Clean build files"
	@echo "  test         - Run tests"
	@echo "  deps         - Download dependencies"
	@echo "  install      - Install binary globally"
	@echo "  examples     - Generate example outputs"
	@echo "  demo         - Run demo conversions"
	@echo "  cross-build  - Build for all platforms"
	@echo "  package      - Create distribution package"