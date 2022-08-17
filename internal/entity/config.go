package entity

import (
	"regexp"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/utils"
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
	Body          bool
	Bot           bool
	Verified      bool
	Report        bool
	Edit          bool
	IgnoredUsers  []string
}

// ReadConfig reads environment variables for input values which are supplied
// from an action runner and create Conventional PR's configuration from it
func ReadConfig() (*Config, error) {
	token := utils.ReadEnvString("INPUT_ACCESS_TOKEN")

	if token == "" {
		return nil, constants.ErrMissingToken
	}

	draft := utils.ReadEnvBool("INPUT_DRAFT")
	close := utils.ReadEnvBool("INPUT_CLOSE")
	strict := utils.ReadEnvBool("INPUT_STRICT")
	assign := utils.ReadEnvBool("INPUT_ASSIGNEE")
	issue := utils.ReadEnvBool("INPUT_ISSUE")
	body := utils.ReadEnvBool("INPUT_BODY")
	bot := utils.ReadEnvBool("INPUT_BOT")
	verified := utils.ReadEnvBool("INPUT_VERIFIED_COMMITS")
	report := utils.ReadEnvBool("INPUT_REPORT")
	edit := utils.ReadEnvBool("INPUT_EDIT")

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

	ignoredUsers := utils.ReadEnvStringArray("INPUT_IGNORED_USERS")

	return &Config{
		Token:         token,
		Draft:         draft,
		Close:         close,
		Edit:          edit,
		Strict:        strict,
		Assign:        assign,
		Issue:         issue,
		TitlePattern:  titlePattern,
		CommitPattern: commitPattern,
		BranchPattern: branchPattern,
		Bot:           bot,
		Label:         label,
		Body:          body,
		Template:      template,
		FileChanges:   fileChanges,
		Verified:      verified,
		IgnoredUsers:  ignoredUsers,
		Report:        report,
	}, nil
}
