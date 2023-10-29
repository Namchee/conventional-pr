package validator

import (
	"context"
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestIssueValidator_IsValid(t *testing.T) {
	type args struct {
		config   bool
		meta     *entity.Meta
		prNumber int
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidationResult
	}{
		{
			name: "should allow issue references",
			args: args{
				prNumber: 1,
				meta: &entity.Meta{
					Name:  "conventional-pr",
					Owner: "Namchee",
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
				prNumber: 2,
				meta: &entity.Meta{
					Name:  "conventional-pr",
					Owner: "Namchee",
				},
				config: false,
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: false,
				Result: nil,
			},
		},
		{
			name: "should reject if no issue references at all",
			args: args{
				prNumber: 2,
				meta: &entity.Meta{
					Name:  "conventional-pr",
					Owner: "Namchee",
				},
				config: true,
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: constants.ErrNoIssue,
			},
		},
		{
			name: "should reject if reference is not on the same repository",
			args: args{
				prNumber: 1,
				meta: &entity.Meta{
					Name:  "conventional-pr",
					Owner: "namcheee",
				},
				config: true,
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: constants.ErrNoIssue,
			},
		},
		{
			name: "should pass if data fetching failed",
			args: args{
				prNumber: 3,
				meta: &entity.Meta{
					Name:  "conventional-pr",
					Owner: "Namchee",
				},
				config: true,
			},
			want: &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config := &entity.Configuration{
				Issue: tc.args.config,
			}
			pullRequest := &entity.PullRequest{
				Number:     tc.args.prNumber,
				Repository: *tc.args.meta,
			}

			client := mocks.NewGithubClientMock()

			validator := NewIssueValidator(client, config)
			got := validator.IsValid(context.TODO(), pullRequest)

			assert.Equal(t, got, tc.want)
		})
	}
}
