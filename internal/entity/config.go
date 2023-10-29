package entity

import (
	"regexp"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/utils"
)

// Configuration is a configuration object that is parsed from the action input
type Configuration struct {
	Token         string
	Draft         bool
	Label         string
	Strict        bool
	Close         bool
	TitlePattern  string
	CommitPattern string
	BranchPattern string
	Message       string
	FileChanges   int
	Issue         bool
	Body          bool
	Bot           bool
	Signed        bool
	Verbose       bool
	Edit          bool
	IgnoredUsers  []string

	RestURL    string
	GraphQLURL string
}

// ReadConfig reads environment variables for input values which are supplied
// from an action runner and create Conventional PR's configuration from it
func ReadConfig() (*Configuration, error) {
	token := utils.ReadEnvString("INPUT_ACCESS_TOKEN")

	if token == "" {
		return nil, constants.ErrMissingToken
	}

	draft := utils.ReadEnvBool("INPUT_DRAFT")
	close := utils.ReadEnvBool("INPUT_CLOSE")
	strict := utils.ReadEnvBool("INPUT_STRICT")
	issue := utils.ReadEnvBool("INPUT_ISSUE")
	body := utils.ReadEnvBool("INPUT_BODY")
	bot := utils.ReadEnvBool("INPUT_BOT")
	signed := utils.ReadEnvBool("INPUT_SIGNED")
	edit := utils.ReadEnvBool("INPUT_EDIT")
	verbose := utils.ReadEnvBool("INPUT_VERBOSE")

	label := utils.ReadEnvString("INPUT_LABEL")
	message := utils.ReadEnvString("INPUT_MESSAGE")

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

	fileChanges := utils.ReadEnvInt("INPUT_MAXIMUM_CHANGES")

	if fileChanges < 0 {
		return nil, constants.ErrNegativeFileChange
	}

	ignoredUsers := utils.ReadEnvStringArray("INPUT_IGNORED_USERS")

	restUrl := utils.ReadEnvString("GITHUB_API_URL")
	graphqlUrl := utils.ReadEnvString("GITHUB_GRAPHQL_URL")

	return &Configuration{
		Token:         token,
		Draft:         draft,
		Close:         close,
		Edit:          edit,
		Strict:        strict,
		Issue:         issue,
		TitlePattern:  titlePattern,
		CommitPattern: commitPattern,
		BranchPattern: branchPattern,
		Bot:           bot,
		Label:         label,
		Body:          body,
		Message:       message,
		FileChanges:   fileChanges,
		Signed:        signed,
		IgnoredUsers:  ignoredUsers,
		Verbose:       verbose,

		RestURL:    restUrl,
		GraphQLURL: graphqlUrl,
	}, nil
}
