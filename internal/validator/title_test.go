package validator

import (
	"testing"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

// This test also tests the default pattern
func TestTitleValidator_IsValid(t *testing.T) {
	type args struct {
		title   string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidationResult
	}{
		{
			name: "should allow valid title",
			args: args{
				title:   "feat: testing",
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.TitleValidatorName,
				Result: nil,
			},
		},
		{
			name: "should skip if pattern is empty",
			args: args{
				title:   "feat: testing",
				pattern: "",
			},
			want: &entity.ValidationResult{
				Name:   constants.TitleValidatorName,
				Result: nil,
			},
		},
		{
			name: "should allow scoped valid title",
			args: args{
				title:   "feat(ci): testing",
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.TitleValidatorName,
				Result: nil,
			},
		},
		{
			name: "should allow breaking changes title",
			args: args{
				title:   "feat(ci)!: testing",
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.TitleValidatorName,
				Result: nil,
			},
		},
		{
			name: "should allow multi line commit message",
			args: args{
				title: `feat(ci): testing
				
				BREAKING CHANGE: foo bar`,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.TitleValidatorName,
				Result: nil,
			},
		},
		{
			name: "should return an error",
			args: args{
				title:   "I'm invalid",
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.TitleValidatorName,
				Result: constants.ErrInvalidTitle,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &github.PullRequest{
				Title: &tc.args.title,
			}
			config := &entity.Config{
				TitlePattern: tc.args.pattern,
			}

			validator := NewTitleValidator(nil, config, nil)

			got := validator.IsValid(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
