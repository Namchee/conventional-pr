package constants

var (
	// Events stores a list of pull request sub-events to be processed
	Events = []string{"opened", "reopened", "ready_for_review", "synchronize", "unlocked", "edited"}
)
