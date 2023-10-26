package validator

import (
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestIsIssueValid(t *testing.T) {
	type args struct {
		config      bool
		meta *entity.Meta
		pullRequest *entity.PullRequest
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidationResult
	}{
		{
			name: "should allow issue references",
			args: args{
				pullRequest: &entity.PullRequest{
					References: []entity.IssueReference{
						{
							Number: 123,
							Meta: entity.Meta{
								Owner: "Namchee",
								Name: "conventional-pr",
							},
						},
					},
				},
				meta: &entity.Meta{
					Name: "conventional-pr",
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
				pullRequest: &entity.PullRequest{},
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
				pullRequest: &entity.PullRequest{
					References: []entity.IssueReference{},
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
			config := &entity.Configuration{
				Issue: tc.args.config,
			}

			validator := NewIssueValidator(config)
			got := validator.IsValid(tc.args.pullRequest)

			assert.Equal(t, got, tc.want)
		})
	}
}
