package entity

import (
	"encoding/json"
	"os"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/utils"
)

// Event that triggers the action
type Event struct {
	// action name
	Action string `json:"action"`
	// pull request number
	Number int `json:"number"`
}

// ReadEvent reads and parse event meta definition
func ReadEvent(path string) (*Event, error) {
	file, err := os.Open(
		utils.ReadEnvString("GITHUB_EVENT_PATH"),
	)

	if err != nil {
		return nil, constants.ErrEventFileRead
	}
	defer file.Close()

	var event Event

	if err = json.NewDecoder(file).Decode(&event); err != nil {
		return nil, constants.ErrEventFileParse
	}

	return &event, nil
}
