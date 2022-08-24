package entity

// ValidationResult represents validator checking result on a pull request
type ValidationResult struct {
	Name   string
	Active bool
	Result error
}

// WhitelistResult represents whitelist checking result on a pull request
type WhitelistResult struct {
	Name   string
	Active bool
	Result bool
}

// PullRequestResult is intermediary type to combine validation and whitelist results
type PullRequestResult struct {
	Validation []*ValidationResult
	Whitelist  []*WhitelistResult
}
