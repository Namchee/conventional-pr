package validator

import (
	"testing"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestBranchValidator_IsValid(t *testing.T) {
	type args struct {
		branch  string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidationResult
	}{
		{
			name: "should allow valid branch name",
			args: args{
				branch:  "feat/namchee",
				pattern: `([\w\-_]+)/[\w_\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.BranchValidatorName,
				Result: nil,
			},
		},
		{
			name: "should skip when pattern is emtpy",
			args: args{
				branch:  "invalid",
				pattern: "",
			},
			want: &entity.ValidationResult{
				Name:   constants.BranchValidatorName,
				Result: nil,
			},
		},
		{
			name: "should return an error",
			args: args{
				branch:  "invalid",
				pattern: `([\w\-_]+)/[\w_\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.BranchValidatorName,
				Result: constants.ErrInvalidBranch,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &github.PullRequest{
				Head: &github.PullRequestBranch{
					Ref: &tc.args.branch,
				},
			}
			config := &entity.Config{
				BranchPattern: tc.args.pattern,
			}

			validator := NewBranchValidator(nil, config, nil)

			got := validator.IsValid(pull)

			assert.Equal(t, tc.want, got)
		})
	}
}