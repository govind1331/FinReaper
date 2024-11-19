package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

// type pageInfo struct {
// 	StatusCode int
// 	Links      map[string]int
// }

// type scrapeData struct {
// 	TextBlob blobstore
// 	Symb string
// }

func main() {

	// Create context with more options for stability
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-background-networking", false),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("headless", false), // Run in visible mode
		chromedp.WindowSize(800, 800),    // Set window size
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create browser context
	ctx, cancel := chromedp.NewContext(allocCtx,
		chromedp.WithLogf(log.Printf), // Enable logging
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	selector := `input[placeholder="Search for news, symbols or companies"]`
	selector_2 := `a[title="View more"]`

	// Run with better error handling and logging
	err := chromedp.Run(ctx,
		// Navigate to the page
		chromedp.Navigate("https://finance.yahoo.com"),

		// Wait for body to be ready
		chromedp.WaitReady("body"),

		// Add a small delay to ensure page is loaded
		chromedp.Sleep(2*time.Second),

		// Log before attempting to find element
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Waiting for search box to become visible...")
			return nil
		}),

		// Wait for the search box
		chromedp.WaitVisible(selector),

		// Log before clicking
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Attempting to click search box...")
			return nil
		}),

		// Click the search box
		chromedp.Click(selector),

		// Ensure element is focused
		chromedp.Focus(selector),

		// Type with small delay between keystrokes
		chromedp.SendKeys(selector, "apple"),

		// Small delay before pressing Enter
		chromedp.Sleep(500*time.Millisecond),

		// Press Enter key
		chromedp.SendKeys(selector, kb.Enter),
		// --------------------------------------------------------------------------------------------------------------
		// Wait to see results
		chromedp.Sleep(10*time.Second),

		// Wait for body to be ready
		// chromedp.WaitReady("body"),

		// Wait for the search box
		chromedp.WaitVisible(selector_2),

		// Click the search box
		chromedp.Click(selector_2),

		// Wait to see results
		chromedp.Sleep(10*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Search completed successfully")

}
