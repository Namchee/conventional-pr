package validator

import (
	"errors"
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestVerifiedValidator_IsValid(t *testing.T) {
	type args struct {
		number   int
		verified bool
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidationResult
	}{
		{
			name: "should allow valid commits",
			args: args{
				number:   123,
				verified: true,
			},
			want: &entity.ValidationResult{
				Name:   constants.VerifiedCommitsValidatorName,
				Result: nil,
			},
		},
		{
			name: "should skip when pattern is empty",
			args: args{
				number:   69,
				verified: false,
			},
			want: &entity.ValidationResult{
				Name:   constants.VerifiedCommitsValidatorName,
				Result: nil,
			},
		},
		{
			name: "should allow when no commits",
			args: args{
				number:   420,
				verified: true,
			},
			want: &entity.ValidationResult{
				Name:   constants.VerifiedCommitsValidatorName,
				Result: nil,
			},
		},
		{
			name: "should reject on invalid commits",
			args: args{
				number:   69,
				verified: true,
			},
			want: &entity.ValidationResult{
				Name:   constants.VerifiedCommitsValidatorName,
				Result: errors.New("commit this is bad is not a verified commit"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &github.PullRequest{
				Number: &tc.args.number,
			}
			config := &entity.Config{
				Verified: tc.args.verified,
			}

			client := mocks.NewGithubClientMock()

			validator := NewVerifiedValidator(client, config, &entity.Meta{})

			got := validator.IsValid(pull)

			assert.Equal(t, tc.want, got)
		})
	}
}
