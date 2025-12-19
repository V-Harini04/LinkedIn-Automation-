package stealth

import (
	"context"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

// SimpleScroll simulates human-like scrolling behavior
func SimpleScroll(ctx context.Context) error {
	rand.Seed(time.Now().UnixNano())

	scrollTimes := rand.Intn(5) + 3 // 3–7 scrolls

	for i := 0; i < scrollTimes; i++ {
		scrollBy := rand.Intn(400) + 200 // scroll 200–600px

		err := chromedp.Run(ctx,
			chromedp.Evaluate(
				`window.scrollBy({ top: `+itoa(scrollBy)+`, behavior: 'smooth' })`,
				nil,
			),
		)
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(rand.Intn(800)+400) * time.Millisecond)
	}

	return nil
}

// helper function (no strconv to keep it simple)
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	digits := ""
	for n > 0 {
		digits = string('0'+n%10) + digits
		n /= 10
	}
	return digits
}
