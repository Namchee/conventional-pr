package whitelist

import (
	"fmt"
	"testing"

	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/mocks"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestBotWhitelist_IsWhitelisted(t *testing.T) {
	type args struct {
		name   string
		config bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be skipped if is bot and bot = true",
			args: args{
				name:   "foo",
				config: true,
			},
			want: true,
		},
		{
			name: "should be checked if is bot and bot = false",
			args: args{
				name:   "foo",
				config: false,
			},
			want: false,
		},
		{
			name: "should be checked if is not bot and bot = true",
			args: args{
				name:   "bar",
				config: true,
			},
			want: false,
		},
		{
			name: "should be checked if is not bot and bot = false",
			args: args{
				name:   "bar",
				config: false,
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			user := &github.User{
				Login: &tc.args.name,
			}
			pull := &github.PullRequest{
				User: user,
			}
			config := &entity.Config{
				Bot: tc.args.config,
			}
			client := mocks.NewGithubClientMock()

			whitelister := NewBotWhitelist(client, config)

			got := whitelister.IsWhitelisted(pull)

			assert.Equal(
				t,
				got,
				tc.want,
				fmt.Sprintf("BotWhitelist.IsWhitelisted() = %v, want = %v", got, tc.want),
			)
		})
	}
}
