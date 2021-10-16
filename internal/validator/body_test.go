package validator

import (
	"testing"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestBodyValidator_IsValid(t *testing.T) {
	type args struct {
		body string
	}

	tests := []struct {
		name string
		args args
		want *entity.ValidatorResult
	}{
		{
			name: "should allow non-empty body",
			args: args{
				body: "foo bar",
			},
			want: &entity.ValidatorResult{
				Name:   constants.BodyValidatorName,
				Result: nil,
			},
		},
		{
			name: "should reject empty body",
			args: args{
				body: "",
			},
			want: &entity.ValidatorResult{
				Name:   constants.BodyValidatorName,
				Result: constants.ErrNoBody,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pullRequest := &github.PullRequest{
				Body: &tc.args.body,
			}

			bodyValidator := NewBodyValidator(nil, nil, nil)

			got := bodyValidator.IsValid(pullRequest)

			assert.Equal(t, got, tc.want)
		})
	}
}
