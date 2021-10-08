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
		want error
	}{
		{
			name: "should allow conventional commit style",
			args: args{
				title:   "feat: testing",
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: nil,
		},
		{
			name: "should allow scoped conventional commit style",
			args: args{
				title:   "feat(ci): testing",
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: nil,
		},
		{
			name: "should allow breaking changes",
			args: args{
				title:   "feat(ci)!: testing",
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: nil,
		},
		{
			name: "should allow multi line commit message",
			args: args{
				title: `feat(ci): testing
				
				BREAKING CHANGE: foo bar`,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: nil,
		},
		{
			name: "should return an error",
			args: args{
				title:   "I'm invalid",
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: constants.ErrInvalidTitle,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &github.PullRequest{
				Title: &tc.args.title,
			}
			config := &entity.Config{
				Pattern: tc.args.pattern,
			}

			validator := NewTitleValidator(nil, config)

			got := validator.IsValid(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
