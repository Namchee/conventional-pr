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
