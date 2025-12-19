package state

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

var (
	stateFile = "state/state.json"
	lock      sync.Mutex
)

type AppState struct {
	LoggedIn        bool      `json:"logged_in"`
	LastLogin       time.Time `json:"last_login"`
	CollectedURLs   []string  `json:"collected_urls"`
	SentConnections []string  `json:"sent_connections"`
	SentMessages    []string  `json:"sent_messages"`
}

func LoadState() (*AppState, error) {
	lock.Lock()
	defer lock.Unlock()

	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		return &AppState{
			CollectedURLs:   []string{},
			SentConnections: []string{},
			SentMessages:    []string{},
		}, nil
	}

	data, err := os.ReadFile(stateFile)
	if err != nil {
		return nil, err
	}

	var state AppState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}

	return &state, nil
}

func SaveState(state *AppState) error {
	lock.Lock()
	defer lock.Unlock()

	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(stateFile, data, 0644)
}
