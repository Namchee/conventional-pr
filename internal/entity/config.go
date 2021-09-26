package entity

import (
	"regexp"

	"github.com/Namchee/ethos/internal/utils"
)

// Configuration object which is parsed from the action input
type Config struct {
	Token         string
	Draft         bool
	Label         string
	Strict        bool
	Close         bool
	Assign        bool
	Template      string
	AllowedTypes  []string
	AllowedScopes []string
	FileChanges   int
	Issue         bool
	Bot           bool
	Commits       bool
}

// ReadConfig reads environment variables for input values which are supplied
// from an action runner and create Conventional PR's configuration from it
func ReadConfig() Config {
	token := utils.ReadEnvString("INPUT_ACCESS_TOKEN")

	if len(token) == 0 {
		// Use github-actions bot account to provide message
		token = utils.ReadEnvString("GITHUB_TOKEN")
	}

	draft := utils.ReadEnvBool("INPUT_CHECK_DRAFT")
	close := utils.ReadEnvBool("INPUT_CLOSE")
	strict := utils.ReadEnvBool("INPUT_STRICT")
	assign := utils.ReadEnvBool("INPUT_ASSIGNEE")
	issue := utils.ReadEnvBool("INPUT_LINK_ISSUE")
	bot := utils.ReadEnvBool("INPUT_IGNORE_BOT")
	commits := utils.ReadEnvBool("INPUT_CHECK_COMMITS")

	label := utils.ReadEnvString("INPUT_LABEL")
	template := utils.ReadEnvString("INPUT_TEMPLATE")

	pattern := regexp.MustCompile(`\s?,\s?`)

	allowedTypes := pattern.Split(
		utils.ReadEnvString("INPUT_ALLOWED_TYPES"),
		-1,
	)
	allowedTypes = utils.RemoveEmptyStrings(allowedTypes)

	allowedScopes := pattern.Split(
		utils.ReadEnvString("INPUT_ALLOWED_SCOPES"),
		-1,
	)
	allowedScopes = utils.RemoveEmptyStrings(allowedScopes)

	fileChanges := utils.ReadEnvInt("INPUT_MAXIMUM_FILE_CHANGES")

	return Config{
		Token:         token,
		Draft:         draft,
		Close:         close,
		Strict:        strict,
		Assign:        assign,
		Issue:         issue,
		Bot:           bot,
		Label:         label,
		Template:      template,
		AllowedTypes:  allowedTypes,
		AllowedScopes: allowedScopes,
		FileChanges:   fileChanges,
		Commits:       commits,
	}
}
