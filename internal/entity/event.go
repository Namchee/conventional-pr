package entity

// Event that triggers the action
type Event struct {
	// action name
	Action string `json:"action"`
	// pull request number
	Number int `json:"number"`
}
