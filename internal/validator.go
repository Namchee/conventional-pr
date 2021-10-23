package internal

import (
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

// Validator validates if a pull request follows a convention or not
type Validator interface {
	IsValid(*github.PullRequest) *entity.ValidationResult
}
