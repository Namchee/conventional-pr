package entity

// ValidationResult represents validator checking result on a pull request
type ValidationResult struct {
	Name   string
	Result error
}

// WhitelistResult represents whitelist checking result on a pull request
type WhitelistResult struct {
	Name   string
	Result bool
}
