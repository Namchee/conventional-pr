package validator

import (
	"context"
	"regexp"
	"strconv"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type issueValidator struct {
	client internal.GithubClient
	config *entity.Config
	meta   *entity.Meta
	Name   string
}

func NewIssueValidator(
	client internal.GithubClient,
	config *entity.Config,
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
	ctx := context.Background()
	pattern := regexp.MustCompile(`#(\d+)`)

	mentions := pattern.FindAllStringSubmatch(pullRequest.GetBody(), -1)

	for _, mention := range mentions {
		num, _ := strconv.Atoi(mention[1])
		issue, err := v.client.GetIssue(ctx, v.meta.Owner, v.meta.Name, num)

		if err == nil && issue != nil {
			return &entity.ValidationResult{
				Name:   v.Name,
				Result: nil,
			}
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Result: constants.ErrNoIssue,
	}
}
