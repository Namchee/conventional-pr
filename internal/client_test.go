package internal

import (
	"testing"

	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewGithubClient(t *testing.T) {
	tests := []struct {
		name    string
		config  *entity.Configuration
		wantErr bool
	}{
		{
			name: "should return nil when GitHub URL is malformed",
			config: &entity.Configuration{
				RestURL: "https://api.github.com:123a",
			},
			wantErr: true,
		},
		{
			name: "should return client when GitHub URL is valid",
			config: &entity.Configuration{
				RestURL: "https://api.github.com/",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewGithubClient(tc.config)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
			}
		})
	}
}
