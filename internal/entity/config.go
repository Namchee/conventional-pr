package entity

import (
	"regexp"

	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/utils"
)

// Config is a configuration object that is parsed from the action input
type Config struct {
	Token         string
	Draft         bool
	Label         string
	Strict        bool
	Close         bool
	Assign        bool
	TitlePattern  string
	CommitPattern string
	BranchPattern string
	Template      string
	FileChanges   int
	Issue         bool
	Bot           bool
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

	label := utils.ReadEnvString("INPUT_LABEL")
	template := utils.ReadEnvString("INPUT_TEMPLATE")

	titlePattern := utils.ReadEnvString("INPUT_TITLE_PATTERN")

	if _, err := regexp.Compile(titlePattern); err != nil {
		return nil, constants.ErrInvalidTitlePattern
	}

	commitPattern := utils.ReadEnvString("INPUT_COMMIT_PATTERN")

	if _, err := regexp.Compile(commitPattern); err != nil {
		return nil, constants.ErrInvalidCommitPattern
	}

	branchPattern := utils.ReadEnvString("INPUT_BRANCH_PATTERN")

	if _, err := regexp.Compile(branchPattern); err != nil {
		return nil, constants.ErrInvalidBranchPattern
	}

	fileChanges := utils.ReadEnvInt("INPUT_MAXIMUM_FILE_CHANGES")

	if fileChanges < 0 {
		return nil, constants.ErrNegativeFileChange
	}

	return &Config{
		Token:         token,
		Draft:         draft,
		Close:         close,
		Strict:        strict,
		Assign:        assign,
		Issue:         issue,
		TitlePattern:  titlePattern,
		CommitPattern: commitPattern,
		BranchPattern: branchPattern,
		Bot:           bot,
		Label:         label,
		Template:      template,
		FileChanges:   fileChanges,
	}, nil
}
