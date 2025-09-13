package gochromedp

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// ConvertOptions holds conversion configuration
type ConvertOptions struct {
	PageSize     string
	Orientation  string
	MarginTop    string
	MarginRight  string
	MarginBottom string
	MarginLeft   string
	Format       string
	Quality      int
	Width        int
	Height       int
	NoBackground bool
	Grayscale    bool
}

// DefaultConvertOptions returns default conversion options
func DefaultConvertOptions() *ConvertOptions {
	return &ConvertOptions{
		PageSize:     "A4",
		Orientation:  "portrait",
		MarginTop:    "10mm",
		MarginRight:  "10mm",
		MarginBottom: "10mm",
		MarginLeft:   "10mm",
		Format:       "png",
		Quality:      90,
		Width:        1024,
		Height:       768,
		NoBackground: false,
		Grayscale:    false,
	}
}

// ConvertHTMLToPDF converts HTML content to PDF
func ConvertHTMLToPDF(htmlContent string, options *ConvertOptions) ([]byte, error) {
	if options == nil {
		options = DefaultConvertOptions()
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create chromedp context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAlloc()

	taskCtx, cancelTask := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancelTask()

	// Navigate to data URL with HTML content
	dataURL := "data:text/html;base64," + base64.StdEncoding.EncodeToString([]byte(htmlContent))

	var pdfData []byte
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(dataURL),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			// Set up PDF print options
			printParams := page.PrintToPDFParams{
				Landscape:         strings.ToLower(options.Orientation) == "landscape",
				PrintBackground:   true,
				PreferCSSPageSize: false,
			}

			// Set page size
			if options.PageSize != "" {
				width, height := getPageDimensions(options.PageSize, strings.ToLower(options.Orientation) == "landscape")
				printParams.PaperWidth = width
				printParams.PaperHeight = height
			}

			// Set margins
			if options.MarginTop != "" {
				margin := parseMarginValue(options.MarginTop)
				printParams.MarginTop = margin
			}
			if options.MarginBottom != "" {
				margin := parseMarginValue(options.MarginBottom)
				printParams.MarginBottom = margin
			}
			if options.MarginLeft != "" {
				margin := parseMarginValue(options.MarginLeft)
				printParams.MarginLeft = margin
			}
			if options.MarginRight != "" {
				margin := parseMarginValue(options.MarginRight)
				printParams.MarginRight = margin
			}

			pdfData, _, err = printParams.Do(ctx)
			return err
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %v", err)
	}

	return pdfData, nil
}

// ConvertURLToPDF converts a URL to PDF
func ConvertURLToPDF(url string, options *ConvertOptions) ([]byte, error) {
	if options == nil {
		options = DefaultConvertOptions()
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create chromedp context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAlloc()

	taskCtx, cancelTask := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancelTask()

	var pdfData []byte
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			// Set up PDF print options
			printParams := page.PrintToPDFParams{
				Landscape:         strings.ToLower(options.Orientation) == "landscape",
				PrintBackground:   true,
				PreferCSSPageSize: false,
			}

			// Set page size
			if options.PageSize != "" {
				width, height := getPageDimensions(options.PageSize, strings.ToLower(options.Orientation) == "landscape")
				printParams.PaperWidth = width
				printParams.PaperHeight = height
			}

			// Set margins
			if options.MarginTop != "" {
				margin := parseMarginValue(options.MarginTop)
				printParams.MarginTop = margin
			}
			if options.MarginBottom != "" {
				margin := parseMarginValue(options.MarginBottom)
				printParams.MarginBottom = margin
			}
			if options.MarginLeft != "" {
				margin := parseMarginValue(options.MarginLeft)
				printParams.MarginLeft = margin
			}
			if options.MarginRight != "" {
				margin := parseMarginValue(options.MarginRight)
				printParams.MarginRight = margin
			}

			pdfData, _, err = printParams.Do(ctx)
			return err
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %v", err)
	}

	return pdfData, nil
}

// ConvertHTMLToImage converts HTML content to image
func ConvertHTMLToImage(htmlContent string, options *ConvertOptions) ([]byte, error) {
	if options == nil {
		options = DefaultConvertOptions()
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create chromedp context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAlloc()

	taskCtx, cancelTask := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancelTask()

	// Navigate to data URL with HTML content
	dataURL := "data:text/html;base64," + base64.StdEncoding.EncodeToString([]byte(htmlContent))

	var imageData []byte
	var fullHeight int64

	// First, navigate and get the full document height
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(dataURL),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Get the full document height
			var height int
			err := chromedp.Evaluate(`Math.max(document.body.scrollHeight, document.documentElement.scrollHeight)`, &height).Do(ctx)
			if err != nil {
				return fmt.Errorf("failed to get document height: %v", err)
			}
			// Add 10% buffer to ensure we capture all content
			fullHeight = int64(float64(height) * 1.1)
			return nil
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get document height: %v", err)
	}

	// Now set viewport to full page height and capture screenshot
	err = chromedp.Run(taskCtx,
		emulation.SetDeviceMetricsOverride(int64(options.Width), fullHeight, 1.0, false),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			imageData, err = page.CaptureScreenshot().Do(ctx)
			return err
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %v", err)
	}

	return imageData, nil
}

// ConvertURLToImage converts a URL to image
func ConvertURLToImage(url string, options *ConvertOptions) ([]byte, error) {
	if options == nil {
		options = DefaultConvertOptions()
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create chromedp context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAlloc()

	taskCtx, cancelTask := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancelTask()

	var imageData []byte
	var fullHeight int64

	// First, navigate and get the full document height
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Get the full document height
			var height int
			err := chromedp.Evaluate(`Math.max(document.body.scrollHeight, document.documentElement.scrollHeight)`, &height).Do(ctx)
			if err != nil {
				return fmt.Errorf("failed to get document height: %v", err)
			}
			// Add 10% buffer to ensure we capture all content
			fullHeight = int64(float64(height) * 1.1)
			return nil
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get document height: %v", err)
	}

	// Now set viewport to full page height and capture screenshot
	err = chromedp.Run(taskCtx,
		emulation.SetDeviceMetricsOverride(int64(options.Width), fullHeight, 1.0, false),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			imageData, err = page.CaptureScreenshot().Do(ctx)
			return err
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to capture screenshot: %v", err)
	}

	return imageData, nil
}

// Helper functions

func getPageDimensions(pageSize string, landscape bool) (float64, float64) {
	// Standard page sizes in inches
	sizes := map[string][2]float64{
		"a4":     {8.27, 11.69},
		"a3":     {11.69, 16.54},
		"letter": {8.5, 10.0},
		"legal":  {8.5, 14.0},
	}

	size, exists := sizes[strings.ToLower(pageSize)]
	if !exists {
		size = sizes["a4"] // default
	}

	if landscape {
		// Swap width and height for landscape
		return size[1], size[0]
	}
	return size[0], size[1]
}

func parseMarginValue(margin string) float64 {
	// Parse margin values (supports mm, cm, in)
	// Convert to inches for chromedp
	if strings.HasSuffix(margin, "mm") {
		if v, err := strconv.ParseFloat(strings.TrimSuffix(margin, "mm"), 64); err == nil {
			return v / 25.4 // mm to inches
		}
	}
	if strings.HasSuffix(margin, "cm") {
		if v, err := strconv.ParseFloat(strings.TrimSuffix(margin, "cm"), 64); err == nil {
			return v / 2.54 // cm to inches
		}
	}
	if strings.HasSuffix(margin, "in") {
		if v, err := strconv.ParseFloat(strings.TrimSuffix(margin, "in"), 64); err == nil {
			return v
		}
	}
	// Try to parse as plain number (assume mm)
	if v, err := strconv.ParseFloat(margin, 64); err == nil {
		return v / 25.4 // assume mm
	}
	// Default to 10mm
	return 10.0 / 25.4
}
