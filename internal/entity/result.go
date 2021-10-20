package entity

type ValidationResult struct {
	Name   string
	Result error
}

type WhitelistResult struct {
	Name   string
	Result bool
}
