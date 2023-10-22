package validator

import (
	"context"
	"regexp"
	"strconv"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

var (
	keywordPattern = regexp.MustCompile(`(?mi)\b(close|closes|closed|fix|fixes|fixed|resolve|resolves|resolved) #(\d+)\b`)
)

type issueValidator struct {
	client internal.GithubClient
	config *entity.Configuration
	meta   *entity.Meta
	Name   string
}

// NewIssueValidator creates a new validator that validates issue resolution
func NewIssueValidator(
	client internal.GithubClient,
	config *entity.Configuration,
	meta *entity.Meta,
) internal.Validator {
	return &issueValidator{
		Name:   constants.IssueValidatorName,
		client: client,
		config: config,
		meta:   meta,
	}
}

func (v *issueValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	if !v.config.Issue {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: false,
			Result: nil,
		}
	}

	keywords := keywordPattern.FindAllStringSubmatch(pullRequest.GetBody(), -1)
	
	if keywords == nil {
		return &entity.ValidationResult{
			Name:   v.Name,
			Active: true,
			Result: constants.ErrNoIssue,
		}
	}

	for _, number := range keywords {
		num, _ := strconv.Atoi(number[2])

		issue, _ := v.client.GetIssue(
			context.Background(),
			v.meta.Owner,
			v.meta.Name,
			num,
		)

		if issue != nil {
			return &entity.ValidationResult{
				Name:   v.Name,
				Active: true,
				Result: nil,
			}
		}
	} 

	return &entity.ValidationResult{
		Name:   v.Name,
		Active: true,
		Result: constants.ErrNoIssue,
	}
}
