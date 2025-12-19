package connect

import (
	"context"
	"log"
	"time"

	"first/state"

	"github.com/chromedp/chromedp"
)

// 🤝 Send connection requests
func SendConnectionRequests(ctx context.Context, limit int) error {
	appState, _ := state.LoadState()
	count := 0

	for _, url := range appState.CollectedURLs {

		if contains(appState.SentConnections, url) {
			log.Println("Already connected — skipping:", url)
			continue
		}

		if count >= limit {
			break
		}

		log.Println("Visiting profile:", url)

		err := chromedp.Run(ctx,
			chromedp.Navigate("about:blank"),
			chromedp.Evaluate(`
				document.body.innerHTML =
				'<div style="margin:50px">' +
				'<h2>Mock Profile</h2>' +
				'<button id="connect">Connect</button>' +
				'</div>';
			`, nil),
		)
		if err != nil {
			log.Println("Navigation failed:", err)
			continue
		}

		time.Sleep(2 * time.Second)

		log.Println("Connection request sent:", url)
		appState.SentConnections = append(appState.SentConnections, url)
		count++
	}

	return state.SaveState(appState)
}

// 💬 Send messages
func SendMessages(ctx context.Context, message string) error {
	appState, _ := state.LoadState()

	for _, url := range appState.SentConnections {

		if contains(appState.SentMessages, url) {
			continue
		}

		log.Println("Messaging:", url)

		err := chromedp.Run(ctx,
			chromedp.Navigate("about:blank"),
			chromedp.Evaluate(`
				document.body.innerHTML =
				'<textarea id="msg"></textarea>';
			`, nil),
		)
		if err != nil {
			log.Println("Failed to send message:", err)
			continue
		}

		appState.SentMessages = append(appState.SentMessages, url)
		time.Sleep(2 * time.Second)
	}

	return state.SaveState(appState)
}

// 🔍 Helper
func contains(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}
