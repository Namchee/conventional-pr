package validator

import (
	"errors"
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/stretchr/testify/assert"
)

// This test also tests the default pattern
func TestIsCommitValid(t *testing.T) {
	type args struct {
		commits []entity.Commit
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidationResult
	}{
		{
			name: "should allow valid commits",
			args: args{
				commits: []entity.Commit{
					{
						Message: "feat: valid commit",
					},
				},
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Active: true,
				Result: nil,
			},
		},
		{
			name: "should skip when pattern is empty",
			args: args{
				commits: []entity.Commit{
					{
						Message: "bad commit",
					},
				},
				pattern: "",
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Active: false,
				Result: nil,
			},
		},
		{
			name: "should allow when no commits",
			args: args{
				commits: []entity.Commit{},
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Active: true,
				Result: nil,
			},
		},
		{
			name: "should reject on invalid commits",
			args: args{
				commits: []entity.Commit{
					{
						Hash: "e21b424",
						Message: "feat: good commit",
					},
					{
						Hash: "e21b423",
						Message: "bad commit",
					},
				},
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Active: true,
				Result: errors.New("commit e21b423 does not have valid commit message"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &entity.PullRequest{
				Commits: tc.args.commits,
			}
			config := &entity.Configuration{
				CommitPattern: tc.args.pattern,
			}

			got := IsCommitValid(pull)

			assert.Equal(t, tc.want, got)
		})
	}
}
