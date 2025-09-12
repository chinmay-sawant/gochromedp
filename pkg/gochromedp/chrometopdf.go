package gochromedp

import (
	"context"
	"fmt"
	"log"
	"net/url"
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
	dataURL := "data:text/html;charset=utf-8," + url.QueryEscape(htmlContent)

	var pdfData []byte
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(dataURL),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfData, _, err = page.PrintToPDF().Do(ctx)
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
			pdfData, _, err = page.PrintToPDF().Do(ctx)
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

	// Set viewport size
	if err := chromedp.Run(taskCtx,
		emulation.SetDeviceMetricsOverride(int64(options.Width), int64(options.Height), 1.0, false),
	); err != nil {
		return nil, fmt.Errorf("failed to set device metrics: %v", err)
	}

	// Navigate to data URL with HTML content
	dataURL := "data:text/html;charset=utf-8," + url.QueryEscape(htmlContent)

	var imageData []byte
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(dataURL),
		chromedp.WaitReady("body"),
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

	// Set viewport size
	if err := chromedp.Run(taskCtx,
		emulation.SetDeviceMetricsOverride(int64(options.Width), int64(options.Height), 1.0, false),
	); err != nil {
		return nil, fmt.Errorf("failed to set device metrics: %v", err)
	}

	var imageData []byte
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
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

func parseDimension(pageSize, orientation string, isWidth bool) float64 {
	// Standard page sizes in inches (converted to inches for chromedp)
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

	if orientation == "landscape" {
		// Swap width and height for landscape
		if isWidth {
			return size[1]
		}
		return size[0]
	}
	if isWidth {
		return size[0]
	}
	return size[1]
}

func parseMargin(margin string) float64 {
	// Simple margin parsing - assumes mm for now
	// Could be extended to support different units
	if strings.HasSuffix(margin, "mm") {
		var value float64
		fmt.Sscanf(margin, "%fmm", &value)
		return value / 25.4 // convert mm to inches
	}
	if strings.HasSuffix(margin, "in") {
		var value float64
		fmt.Sscanf(margin, "%fin", &value)
		return value
	}
	// Default to 10mm
	return 10.0 / 25.4
}
