package stealth

import (
	"context"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

// ApplyFingerprintMask applies basic browser fingerprint masking
func ApplyFingerprintMask(ctx context.Context) error {
	return chromedp.Run(ctx,

		// 1️⃣ Set realistic viewport
		emulation.SetDeviceMetricsOverride(
			1366,
			768,
			1,
			false,
		),

		// 2️⃣ Set common User-Agent
		emulation.SetUserAgentOverride(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
				"AppleWebKit/537.36 (KHTML, like Gecko) "+
				"Chrome/120.0.0.0 Safari/537.36",
		),

		// 3️⃣ Remove webdriver flag
		chromedp.Evaluate(`
			Object.defineProperty(navigator, 'webdriver', {
				get: () => undefined
			});
		`, nil),

		// 4️⃣ Fake plugins & languages
		chromedp.Evaluate(`
			Object.defineProperty(navigator, 'plugins', {
				get: () => [1, 2, 3, 4, 5],
			});

			Object.defineProperty(navigator, 'languages', {
				get: () => ['en-US', 'en'],
			});
		`, nil),
	)
}
