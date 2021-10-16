package validator

import (
	"errors"
	"testing"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

// This test also tests the default pattern
func TestCommitValidator_IsValid(t *testing.T) {
	type args struct {
		number  int
		config  bool
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidatorResult
	}{
		{
			name: "should allow valid commits",
			args: args{
				number:  123,
				config:  true,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidatorResult{
				Name:   constants.CommitValidatorName,
				Result: nil,
			},
		},
		{
			name: "should allow when config is false",
			args: args{
				number:  69,
				config:  false,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidatorResult{
				Name:   constants.CommitValidatorName,
				Result: nil,
			},
		},
		{
			name: "should allow when no commits",
			args: args{
				number:  420,
				config:  true,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidatorResult{
				Name:   constants.CommitValidatorName,
				Result: nil,
			},
		},
		{
			name: "should reject on invalid commits",
			args: args{
				number:  69,
				config:  true,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidatorResult{
				Name:   constants.CommitValidatorName,
				Result: errors.New("commit this is bad does not have valid commit message"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &github.PullRequest{
				Number: &tc.args.number,
			}
			config := &entity.Config{
				Commits: tc.args.config,
				Pattern: tc.args.pattern,
			}

			client := mocks.NewGithubClientMock()

			validator := NewCommitValidator(client, config, &entity.Meta{})

			got := validator.IsValid(pull)

			assert.Equal(t, tc.want, got)
		})
	}
}
