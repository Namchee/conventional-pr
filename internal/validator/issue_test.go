package validator

import (
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestIssueValidator_IsValid(t *testing.T) {
	type args struct {
		config      bool
		pullRequest *github.PullRequest
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidationResult
	}{
		{
			name: "should allow issue references",
			args: args{
				pullRequest: &github.PullRequest{
					Body: github.String(`
						## Overview

						Closes #1

						This pull request resolves #123 by using xxx/yyy/zzz.
					`),
				},
				config: true,
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: nil,
			},
		},
		{
			name: "should be skipped if disabled",
			args: args{
				pullRequest: &github.PullRequest{},
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: false,
				Result: nil,
			},
		},
		{
			name: "should reject if no issue at all",
			args: args{
				pullRequest: &github.PullRequest{
					Body: github.String("I don't reference any issues!"),
				},
				config:      true,
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: constants.ErrNoIssue,
			},
		},
		{
			name: "should reject if references are fake",
			args: args{
				pullRequest: &github.PullRequest{
					Body: github.String(`
						## Overview

						Closes #1. resolved #2. fixes #3. Fix #321Closed #123

						This pull request resolve #124 by using xxx/yyy/zzz.






						close #9. Fixed #54.
					`),
				},
				config:      true,
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: constants.ErrNoIssue,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := mocks.NewGithubClientMock()
			config := &entity.Configuration{
				Issue: tc.args.config,
			}

			validator := NewIssueValidator(client, config, &entity.Meta{})

			got := validator.IsValid(tc.args.pullRequest)

			assert.Equal(t, got, tc.want)
		})
	}
}
