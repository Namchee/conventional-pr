package validator

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestValidatorGroup(t *testing.T) {
	pullRequest := &entity.PullRequest{
		Number: 123,
	}

	config := &entity.Configuration{}

	wg := sync.WaitGroup{}
	client := mocks.NewGithubClientMock()

	validatorGroup := NewValidatorGroup(
		client,
		config,
		&wg,
	)

	got := validatorGroup.Process(context.TODO(), pullRequest)

	assert.Equal(t, 6, len(got))
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		args []*entity.ValidationResult
		want bool
	}{
		{
			name: "should return true",
			args: []*entity.ValidationResult{
				{
					Name:   "foo bar",
					Active: true,
					Result: nil,
				},
				{
					Name:   "bar baz",
					Active: true,
					Result: nil,
				},
			},
			want: true,
		},
		{
			name: "should return false",
			args: []*entity.ValidationResult{
				{
					Name:   "foo bar",
					Active: true,
					Result: nil,
				},
				{
					Name:   "bar baz",
					Active: true,
					Result: errors.New("testing"),
				},
			},
			want: false,
		},
		{
			name: "should ignore error when inactive",
			args: []*entity.ValidationResult{
				{
					Name:   "foo bar",
					Active: false,
					Result: errors.New("foo bar"),
				},
				{
					Name:   "bar baz",
					Active: true,
					Result: nil,
				},
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.args)

			assert.Equal(t, tc.want, got)
		})
	}
}
