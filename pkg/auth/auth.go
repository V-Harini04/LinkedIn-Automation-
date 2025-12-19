package auth

import (
	"log"
	"time"

	"first/state"
)

func Authenticate() error {
	appState, err := state.LoadState()
	if err != nil {
		return err
	}

	// Already logged in recently
	if appState.LoggedIn && time.Since(appState.LastLogin) < 24*time.Hour {
		log.Println("Already authenticated — skipping login")
		return nil
	}

	// Simulate login
	log.Println("Performing authentication...")
	time.Sleep(2 * time.Second)

	appState.LoggedIn = true
	appState.LastLogin = time.Now()

	return state.SaveState(appState)
}
