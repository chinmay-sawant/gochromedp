package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	gochromedp "github.com/chinmay-sawant/gochromedp/pkg/gochromedp"
	"github.com/spf13/cobra"
)

var (
	version      = "1.0.0"
	pageSize     string
	orientation  string
	marginTop    string
	marginRight  string
	marginBottom string
	marginLeft   string
	format       string
	quality      int
	width        int
	height       int
)

var rootCmd = &cobra.Command{
	Use:   "gochromedp",
	Short: "Convert HTML to PDF/Image using Chromium headless",
	Long: `gochromedp is a command-line tool similar to wkhtmltopdf but uses Chromium headless
for better rendering and modern web standards support.

Examples:
  # Convert URL to PDF
  gochromedp https://example.com output.pdf

  # Convert HTML file to PDF with custom options
  gochromedp --page-size A4 --margin-top 20mm input.html output.pdf

  # Convert URL to PNG image
  gochromedp --format png --width 1024 --height 768 https://example.com screenshot.png`,
}

var pdfCmd = &cobra.Command{
	Use:   "pdf [input] [output]",
	Short: "Convert HTML to PDF",
	Long:  `Convert HTML content from a file or URL to PDF format.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		output := args[1]

		if err := convertToPDF(input, output); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("✅ Successfully converted %s to %s\n", input, output)
	},
}

var imageCmd = &cobra.Command{
	Use:   "image [input] [output]",
	Short: "Convert HTML to image",
	Long:  `Convert HTML content from a file or URL to image format (PNG/JPEG).`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		output := args[1]

		if err := convertToImage(input, output); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("✅ Successfully converted %s to %s\n", input, output)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gochromedp version %s\n", version)
	},
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&pageSize, "page-size", "A4", "Page size (A4, A3, Letter, etc.)")
	rootCmd.PersistentFlags().StringVar(&orientation, "orientation", "portrait", "Page orientation (portrait/landscape)")
	rootCmd.PersistentFlags().StringVar(&marginTop, "margin-top", "10mm", "Top margin")
	rootCmd.PersistentFlags().StringVar(&marginRight, "margin-right", "10mm", "Right margin")
	rootCmd.PersistentFlags().StringVar(&marginBottom, "margin-bottom", "10mm", "Bottom margin")
	rootCmd.PersistentFlags().StringVar(&marginLeft, "margin-left", "10mm", "Left margin")

	// PDF specific flags
	pdfCmd.Flags().Bool("no-background", false, "Do not print background")
	pdfCmd.Flags().Bool("grayscale", false, "Generate grayscale PDF")

	// Image specific flags
	imageCmd.Flags().StringVar(&format, "format", "png", "Image format (png/jpeg)")
	imageCmd.Flags().IntVar(&quality, "quality", 90, "Image quality (1-100, for JPEG)")
	imageCmd.Flags().IntVar(&width, "width", 1024, "Viewport width")
	imageCmd.Flags().IntVar(&height, "height", 768, "Viewport height")

	rootCmd.AddCommand(pdfCmd)
	rootCmd.AddCommand(imageCmd)
	rootCmd.AddCommand(versionCmd)
}

func convertToPDF(input, output string) error {
	options := &gochromedp.ConvertOptions{
		PageSize:     pageSize,
		Orientation:  orientation,
		MarginTop:    marginTop,
		MarginRight:  marginRight,
		MarginBottom: marginBottom,
		MarginLeft:   marginLeft,
	}

	var data []byte
	var err error

	// Check if input is a URL or file
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		data, err = gochromedp.ConvertURLToPDF(input, options)
	} else {
		// Read HTML file
		htmlContent, readErr := os.ReadFile(input)
		if readErr != nil {
			return fmt.Errorf("failed to read input file: %v", readErr)
		}
		data, err = gochromedp.ConvertHTMLToPDF(string(htmlContent), options)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(output, data, 0644)
}

func convertToImage(input, output string) error {
	options := &gochromedp.ConvertOptions{
		Format:  format,
		Quality: quality,
		Width:   width,
		Height:  height,
	}

	var data []byte
	var err error

	// Check if input is a URL or file
	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		data, err = gochromedp.ConvertURLToImage(input, options)
	} else {
		// Read HTML file
		htmlContent, readErr := os.ReadFile(input)
		if readErr != nil {
			return fmt.Errorf("failed to read input file: %v", readErr)
		}
		data, err = gochromedp.ConvertHTMLToImage(string(htmlContent), options)
	}

	if err != nil {
		return err
	}

	return os.WriteFile(output, data, 0644)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
