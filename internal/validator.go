package internal

import (
	"context"

	"github.com/Namchee/conventional-pr/internal/entity"
)

// Validator validates if a pull request follows a convention or not
type Validator interface {
	IsValid(context.Context, *entity.PullRequest) *entity.ValidationResult
}
