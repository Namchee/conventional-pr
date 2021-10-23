package internal

import (
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

// Whitelist checks if a pull request validation may be skipped or not
type Whitelist interface {
	IsWhitelisted(*github.PullRequest) *entity.WhitelistResult
}
