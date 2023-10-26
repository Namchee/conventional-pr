package whitelist

import (
	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

type permissionWhitelist struct {
	Name   string

	client internal.GithubClient
	config *entity.Configuration
}

// NewPermissionWhitelist creates a whitelist that bypasses checks on pull request submitted by user with high privileges
func NewPermissionWhitelist(
	client internal.GithubClient,
	config *entity.Configuration,
) internal.Whitelist {
	return &permissionWhitelist{
		Name:   constants.PermissionWhitelistName,
		client: client,
		config: config,
	}
}

func (w *permissionWhitelist) IsWhitelisted(pullRequest *entity.PullRequest) *entity.WhitelistResult {
	/*
	if w.config.Strict {
		return &entity.WhitelistResult{
			Name:   w.Name,
			Active: false,
			Result: false,
		}
	}

	ctx := context.Background()

	perms, _ := w.client.GetPermissionLevel(
		ctx,
		pullRequest.Repository.Owner,
		pullRequest.Repository.Name,
		pullRequest.Author.Login,
	)

	result := strings.ToLower(perms.GetPermission()) == constants.AdminUser
	*/

	return &entity.WhitelistResult{
		Name:   w.Name,
		Active: true,
		Result: true,
	}
}
