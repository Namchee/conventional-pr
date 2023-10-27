package validator

import (
	"context"
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestBodyValidator_IsValid(t *testing.T) {
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
			name: "should allow non-empty body",
			args: args{
				config: true,
				body:   "foo bar",
			},
			want: &entity.ValidationResult{
				Name:   constants.BodyValidatorName,
				Active: true,
				Result: nil,
			},
		},
		{
			name: "should be skipped when disabled",
			args: args{
				body: "",
			},
			want: &entity.ValidationResult{
				Name:   constants.BodyValidatorName,
				Active: false,
				Result: nil,
			},
		},
		{
			name: "should reject empty body",
			args: args{
				config: true,
				body:   "",
			},
			want: &entity.ValidationResult{
				Name:   constants.BodyValidatorName,
				Active: true,
				Result: constants.ErrNoBody,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pullRequest := &entity.PullRequest{
				Body: tc.args.body,
			}
			config := &entity.Configuration{
				Body: tc.args.config,
			}

			validator := NewBodyValidator(nil, config)
			got := validator.IsValid(context.TODO(), pullRequest)

			assert.Equal(t, got, tc.want)
		})
	}
}
