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
				number:  123,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Result: nil,
			},
		},
		{
			name: "should skip when pattern is empty",
			args: args{
				number:  123,
				pattern: "",
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Result: nil,
			},
		},
		{
			name: "should allow when no commits",
			args: args{
				number:  420,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Result: nil,
			},
		},
		{
			name: "should reject on invalid commits",
			args: args{
				number:  69,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
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
				CommitPattern: tc.args.pattern,
			}

			client := mocks.NewGithubClientMock()

			validator := NewCommitValidator(client, config, &entity.Meta{})

			got := validator.IsValid(pull)

			assert.Equal(t, tc.want, got)
		})
	}
}
