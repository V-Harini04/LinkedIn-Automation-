package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"first/pkg/auth"
	"first/pkg/connect"
	"first/pkg/search"
	"first/pkg/stealth"
	"first/state"

	"github.com/chromedp/chromedp"
)

func main() {

	// 🔐 Authentication
	if err := auth.Authenticate(); err != nil {
		log.Fatal("Authentication failed:", err)
	}

	log.Println("Starting browser automation...")
	rand.Seed(time.Now().UnixNano())

	// 🧠 Chrome allocator
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer allocCancel()

	// 🌐 Browser context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// ⏱ Timeout
	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	// 🌍 Open page
	if err := chromedp.Run(ctx, chromedp.Navigate("about:blank")); err != nil {
		log.Fatal(err)
	}

	// 🕵 Fingerprint masking
	if err := stealth.ApplyFingerprintMask(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Browser fingerprint masking applied")

	// 🔍 Search
	results, err := search.SearchAndCollect(ctx, "golang developer linkedin")
	if err != nil {
		log.Fatal(err)
	}

	// 💾 Save URLs
	appState, err := state.LoadState()
	if err != nil {
		log.Fatal("Failed to load state:", err)
	}
	appState.CollectedURLs = uniqueStrings(
		append(appState.CollectedURLs, results...),
	)
	if err := state.SaveState(appState); err != nil {
		log.Fatal("Failed to save state:", err)
	}
	// 🤝 Send connections
	if err := connect.SendConnectionRequests(ctx, 3); err != nil {
		log.Println("Connection error:", err)
	}

	// 💬 Send messages
	if err := connect.SendMessages(ctx, "Hi, I’d love to connect with you!"); err != nil {
		log.Println("Message error:", err)
	}

	// 🧪 Demo input
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(`
			if (!document.getElementById("test-input")) {
				const el = document.createElement("input");
				el.id = "test-input";
				el.style.margin = "40px";
				document.body.appendChild(el);
			}
		`, nil),
	); err != nil {
		log.Fatal(err)
	}

	// 🖱 Mouse
	_ = stealth.HumanMouseMove(ctx)
	_ = stealth.SimpleScroll(ctx)

	// ⌨ Typing
	_ = stealth.HumanType(ctx, "#test-input", "Hello, this is human typing")

	log.Println("END OF MAIN REACHED — EXITING")
}
func uniqueStrings(input []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
