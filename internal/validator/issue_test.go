package validator

import (
	"testing"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestIssueValidator_IsValid(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name string
		args args
		want *entity.ValidatorResult
	}{
		{
			name: "should allow issue references",
			args: args{
				body: "Closes #123",
			},
			want: &entity.ValidatorResult{
				Name:   constants.IssueValidatorName,
				Result: nil,
			},
		},
		{
			name: "should reject if no issue at all",
			args: args{
				body: "this is a fake body",
			},
			want: &entity.ValidatorResult{
				Name:   constants.IssueValidatorName,
				Result: constants.ErrNoIssue,
			},
		},
		{
			name: "should distinguih false alarm",
			args: args{
				body: "This is a fake issue #69",
			},
			want: &entity.ValidatorResult{
				Name:   constants.IssueValidatorName,
				Result: constants.ErrNoIssue,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pull := &github.PullRequest{
				Body: &tc.args.body,
			}

			client := mocks.NewGithubClientMock()

			validator := NewIssueValidator(client, nil, &entity.Meta{})

			got := validator.IsValid(pull)

			assert.Equal(t, got, tc.want)
		})
	}
}
