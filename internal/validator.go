package internal

import (
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type Validator interface {
	IsValid(*github.PullRequest) *entity.ValidationResult
}
