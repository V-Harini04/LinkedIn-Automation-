package search

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func SearchAndCollect(ctx context.Context, query string) ([]string, error) {

	log.Println("Starting search for:", query)

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.google.com"),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return nil, err
	}

	// ⚠️ Demo result (safe placeholder)
	results := []string{
		"https://linkedin.com/in/sample-profile-3",
		"https://linkedin.com/in/sample-profile-4",
	}

	log.Printf("Collected %d profile URLs\n", len(results))
	return results, nil
}
