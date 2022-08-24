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
		config bool
		body   string
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidationResult
	}{
		{
			name: "should allow issue references",
			args: args{
				body:   "Closes #123",
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
				body: "Closes #123",
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
				body:   "this is a fake body",
				config: true,
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: constants.ErrNoIssue,
			},
		},
		{
			name: "should distinguih false alarm",
			args: args{
				body:   "This is a fake issue #69",
				config: true,
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
			pull := &github.PullRequest{
				Body: &tc.args.body,
			}

			client := mocks.NewGithubClientMock()
			config := &entity.Configuration{
				Issue: tc.args.config,
			}

			validator := NewIssueValidator(client, config, &entity.Meta{})

			got := validator.IsValid(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
