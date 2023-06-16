package formatter

import (
	"errors"
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
		want string
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
			want: `
 ____                            _   _                   _ ____  ____  
/ ___| ___  _ ____   _____ _ __ | |_(_) ___  _ __   __ _| |  _ \|  _ \ 
| |   / _ \| '_ \ \ / / _ \ '_ \| __| |/ _ \| '_ \ / _ || | |_) | |_) |
| |__| (_) | | | \ V /  __/ | | | |_| | (_) | | | | (_||| |  __/|  _ < 
\____ \___/|_| |_|\_/ \___|_| |_|\__|_|\___/|_| |_|\__,_|_|_|   |_| \_\

┌─────────────────────────────┐
│ Whitelist Result            │
├───────────┬────────┬────────┤
│ WHITELIST │ ACTIVE │ RESULT │
├───────────┼────────┼────────┤
│ foo bar   │ ✅     │ ✅     │
└───────────┴────────┴────────┘

Result: Pull request matches with one (or more) enabled whitelist criteria. Pull request validation is skipped.`,
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
			want: `
 ____                            _   _                   _ ____  ____  
/ ___| ___  _ ____   _____ _ __ | |_(_) ___  _ __   __ _| |  _ \|  _ \ 
| |   / _ \| '_ \ \ / / _ \ '_ \| __| |/ _ \| '_ \ / _ || | |_) | |_) |
| |__| (_) | | | \ V /  __/ | | | |_| | (_) | | | | (_||| |  __/|  _ < 
\____ \___/|_| |_|\_/ \___|_| |_|\__|_|\___/|_| |_|\__,_|_|_|   |_| \_\

┌─────────────────────────────┐
│ Whitelist Result            │
├───────────┬────────┬────────┤
│ WHITELIST │ ACTIVE │ RESULT │
├───────────┼────────┼────────┤
│ foo bar   │ ✅     │ ❌     │
└───────────┴────────┴────────┘

Result: Pull request does not satisfy any enabled whitelist criteria. Pull request will be validated.

┌──────────────────────────────┐
│ Validation Result            │
├────────────┬────────┬────────┤
│ VALIDATION │ ACTIVE │ RESULT │
├────────────┼────────┼────────┤
│ bar baz    │ ✅     │ ✅     │
└────────────┴────────┴────────┘

Result: Pull request satisfies all enabled pull request rules.`,
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
			want: `
 ____                            _   _                   _ ____  ____  
/ ___| ___  _ ____   _____ _ __ | |_(_) ___  _ __   __ _| |  _ \|  _ \ 
| |   / _ \| '_ \ \ / / _ \ '_ \| __| |/ _ \| '_ \ / _ || | |_) | |_) |
| |__| (_) | | | \ V /  __/ | | | |_| | (_) | | | | (_||| |  __/|  _ < 
\____ \___/|_| |_|\_/ \___|_| |_|\__|_|\___/|_| |_|\__,_|_|_|   |_| \_\

┌─────────────────────────────┐
│ Whitelist Result            │
├───────────┬────────┬────────┤
│ WHITELIST │ ACTIVE │ RESULT │
├───────────┼────────┼────────┤
│ foo bar   │ ✅     │ ❌     │
└───────────┴────────┴────────┘

Result: Pull request does not satisfy any enabled whitelist criteria. Pull request will be validated.

┌──────────────────────────────┐
│ Validation Result            │
├────────────┬────────┬────────┤
│ VALIDATION │ ACTIVE │ RESULT │
├────────────┼────────┼────────┤
│ bar baz    │ ✅     │ ❌     │
└────────────┴────────┴────────┘

Result: Pull request is invalid.

Reason:
- Testing`,
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
						Name:   "fuuu",
						Active: true,
						Result: errors.New("This is testing"),
					},
					{
						Name:   "def",
						Active: false,
						Result: nil,
					},
				},
			},
			want: `
 ____                            _   _                   _ ____  ____  
/ ___| ___  _ ____   _____ _ __ | |_(_) ___  _ __   __ _| |  _ \|  _ \ 
| |   / _ \| '_ \ \ / / _ \ '_ \| __| |/ _ \| '_ \ / _ || | |_) | |_) |
| |__| (_) | | | \ V /  __/ | | | |_| | (_) | | | | (_||| |  __/|  _ < 
\____ \___/|_| |_|\_/ \___|_| |_|\__|_|\___/|_| |_|\__,_|_|_|   |_| \_\

┌─────────────────────────────┐
│ Whitelist Result            │
├───────────┬────────┬────────┤
│ WHITELIST │ ACTIVE │ RESULT │
├───────────┼────────┼────────┤
│ foo bar   │ ✅     │ ❌     │
│ abc       │ ❌     │ ❌     │
└───────────┴────────┴────────┘

Result: Pull request does not satisfy any enabled whitelist criteria. Pull request will be validated.

┌──────────────────────────────┐
│ Validation Result            │
├────────────┬────────┬────────┤
│ VALIDATION │ ACTIVE │ RESULT │
├────────────┼────────┼────────┤
│ bar baz    │ ✅     │ ❌     │
│ fuuu       │ ✅     │ ❌     │
│ def        │ ❌     │ ✅     │
└────────────┴────────┴────────┘

Result: Pull request is invalid.

Reason:
- Testing
- This is testing`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			results := &entity.PullRequestResult{
				Whitelist:  tc.args.whitelist,
				Validation: tc.args.validation,
			}

			got := FormatResultToConsole(results)

			assert.Equal(t, tc.want, got)
		})
	}
}
