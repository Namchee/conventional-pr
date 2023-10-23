package repository

import (
	"testing"

	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewGithubRESTClient(t *testing.T) {
	assert.NotPanics(t, func() {
		config := &entity.Configuration{
			Token: "abcde",
		}

		NewGithubRESTClient(config)
	})
}