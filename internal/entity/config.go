package entity

import (
	"regexp"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/utils"
)

// Config is a configuration object that is parsed from the action input
type Config struct {
	Token       string
	Draft       bool
	Label       string
	Strict      bool
	Close       bool
	Assign      bool
	Pattern     string
	Template    string
	FileChanges int
	Issue       bool
	Bot         bool
	Commits     bool
}

// ReadConfig reads environment variables for input values which are supplied
// from an action runner and create Conventional PR's configuration from it
func ReadConfig() (*Config, error) {
	token := utils.ReadEnvString("INPUT_ACCESS_TOKEN")

	if token == "" {
		return nil, constants.ErrMissingToken
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

	pattern := utils.ReadEnvString("INPUT_PATTERN")

	if _, err := regexp.Compile(pattern); err != nil {
		return nil, constants.ErrInvalidTitle
	}

	fileChanges := utils.ReadEnvInt("INPUT_MAXIMUM_FILE_CHANGES")

	if fileChanges < 0 {
		return nil, constants.ErrNegativeFileChange
	}

	return &Config{
		Token:       token,
		Draft:       draft,
		Close:       close,
		Strict:      strict,
		Assign:      assign,
		Issue:       issue,
		Pattern:     pattern,
		Bot:         bot,
		Label:       label,
		Template:    template,
		FileChanges: fileChanges,
		Commits:     commits,
	}, nil
}
