package internal

import (
	"github.com/google/go-github/v32/github"
)

type Validator interface {
	IsValid(*github.PullRequest) error
}

type Whitelist interface {
	IsWhitelisted(*github.PullRequest) bool
}
