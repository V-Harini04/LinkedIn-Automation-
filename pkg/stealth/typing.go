package stealth

import (
	"context"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

// HumanType types text character by character with random delay
func HumanType(ctx context.Context, selector string, text string) error {
	// Ensure element is focused
	if err := chromedp.Run(ctx,
		chromedp.Focus(selector),
	); err != nil {
		return err
	}

	for _, char := range text {
		delay := time.Duration(50+rand.Intn(150)) * time.Millisecond

		err := chromedp.Run(ctx,
			chromedp.SendKeys(selector, string(char)),
		)
		if err != nil {
			return err
		}

		time.Sleep(delay)
	}

	return nil
}
