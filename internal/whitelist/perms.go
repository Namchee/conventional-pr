package whitelist

import (
	"context"
	"strings"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type permissionWhitelist struct {
	client internal.GithubClient
	config *entity.Config
	meta   *entity.Meta
	Name   string
}

func NewPermissionWhitelist(
	client internal.GithubClient,
	config *entity.Config,
	meta *entity.Meta,
) internal.Whitelist {
	return &permissionWhitelist{
		Name:   constants.PermissionWhitelistName,
		client: client,
		config: config,
		meta:   meta,
	}
}

func (w *permissionWhitelist) IsWhitelisted(pullRequest *github.PullRequest) *entity.WhitelistResult {
	ctx := context.Background()

	perms, _ := w.client.GetPermissionLevel(
		ctx,
		w.meta.Owner,
		w.meta.Name,
		pullRequest.GetUser().GetLogin(),
	)

	result := strings.ToLower(perms.GetPermission()) == constants.AdminUser &&
		!w.config.Strict

	return &entity.WhitelistResult{
		Name:   w.Name,
		Result: result,
	}
}
