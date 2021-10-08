package validator

import (
	"testing"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestFileValidator_IsValid(t *testing.T) {
	type args struct {
		changes int
		config  int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "should allow few changes",
			args: args{
				changes: 2,
				config:  2,
			},
			want: nil,
		},
		{
			name: "should reject if introduces too many change",
			args: args{
				changes: 3,
				config:  2,
			},
			want: constants.ErrTooManyChanges,
		},
		{
			name: "should allow huge changes if turned off",
			args: args{
				changes: 10000,
				config:  0,
			},
			want: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &github.PullRequest{
				ChangedFiles: &tc.args.changes,
			}
			config := &entity.Config{
				FileChanges: tc.args.config,
			}

			validator := NewFileValidator(nil, config, nil)

			got := validator.IsValid(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
