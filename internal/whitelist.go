package internal

import (
	"github.com/Namchee/conventional-pr/internal/entity"
)

// Whitelist checks if a pull request validation may be skipped or not
type Whitelist interface {
	IsWhitelisted(*entity.PullRequest) *entity.WhitelistResult
}
