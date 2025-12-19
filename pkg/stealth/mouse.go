package stealth

import (
	"context"
	"math/rand"
	"time"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
)

// HumanMouseMove simulates human-like mouse movement
func HumanMouseMove(ctx context.Context) error {
	rand.Seed(time.Now().UnixNano())

	// Get viewport size
	var width, height float64
	err := chromedp.Run(ctx,
		chromedp.Evaluate(`window.innerWidth`, &width),
		chromedp.Evaluate(`window.innerHeight`, &height),
	)
	if err != nil {
		return err
	}

	// Initial cursor position
	x := rand.Float64() * width
	y := rand.Float64() * height

	steps := rand.Intn(20) + 20 // 20–40 steps

	for i := 0; i < steps; i++ {
		x += rand.Float64()*30 - 15
		y += rand.Float64()*30 - 15

		// Clamp inside viewport
		if x < 0 {
			x = 0
		}
		if y < 0 {
			y = 0
		}
		if x > width {
			x = width
		}
		if y > height {
			y = height
		}

		err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
			return input.DispatchMouseEvent(
				input.MouseMoved, // CORRECT constant
				x,
				y,
			).Do(ctx)
		}))
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(rand.Intn(30)+10) * time.Millisecond)
	}

	return nil
}
