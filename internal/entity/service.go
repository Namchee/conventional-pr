package entity

type ValidatorResult struct {
	Name   string
	Result error
}

type WhitelistResult struct {
	Name   string
	Result bool
}
