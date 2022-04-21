package formatter

import (
	"errors"
	"log"
	"testing"

	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestFormatResultToConsole(t *testing.T) {
	type args struct {
		whitelist  []*entity.WhitelistResult
		validation []*entity.ValidationResult
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "formats whitelisted pull request correctly",
			args: args{
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo bar",
						Active: true,
						Result: true,
					},
				},
				validation: []*entity.ValidationResult{},
			},
		},
		{
			name: "formats valid pull request correctly",
			args: args{
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo bar",
						Active: true,
						Result: false,
					},
				},
				validation: []*entity.ValidationResult{
					{
						Name:   "bar baz",
						Active: true,
						Result: nil,
					},
				},
			},
		},
		{
			name: "format invalid pull request correctly",
			args: args{
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo bar",
						Active: true,
						Result: false,
					},
				},
				validation: []*entity.ValidationResult{
					{
						Name:   "bar baz",
						Active: true,
						Result: errors.New("testing"),
					},
				},
			},
		},
		{
			name: "able to format inactive whitelist and validator",
			args: args{
				whitelist: []*entity.WhitelistResult{
					{
						Name:   "foo bar",
						Active: true,
						Result: false,
					},
					{
						Name:   "abc",
						Active: false,
						Result: false,
					},
				},
				validation: []*entity.ValidationResult{
					{
						Name:   "bar baz",
						Active: true,
						Result: errors.New("testing"),
					},
					{
						Name:   "def",
						Active: false,
						Result: nil,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.Default()

			assert.NotPanics(t, func() {
				FormatResultToConsole(tc.args.whitelist, tc.args.validation, logger)
			})
		})
	}
}
