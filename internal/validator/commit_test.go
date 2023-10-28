package validator

import (
	"context"
	"errors"
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
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
				Active: true,
				Result: nil,
			},
		},
		{
			name: "should skip when pattern is empty",
			args: args{
				number:  69,
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
				number:  456,
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
				number:  69,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Active: true,
				Result: errors.New("commit e21b424 does not have valid commit message"),
			},
		},
		{
			name: "should pass when fetch fails",
			args: args{
				number:  1,
				pattern: `([\w\-]+)(\([\w\-]+\))?!?: [\w\s:\-]+`,
			},
			want: &entity.ValidationResult{
				Name:   constants.CommitValidatorName,
				Active: true,
				Result: nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &entity.PullRequest{
				Number: tc.args.number,
			}
			config := &entity.Configuration{
				CommitPattern: tc.args.pattern,
			}

			client := mocks.NewGithubClientMock()

			validator := NewCommitValidator(client, config)
			got := validator.IsValid(context.TODO(), pull)

			assert.Equal(t, tc.want, got)
		})
	}
}
