package constants

import "errors"

// Metadata error
var (
	ErrMalformedMetadata = errors.New("[Meta] Malformed repository metadata")
)

// Config error
var (
	ErrMissingToken         = errors.New("[Config] Access token is empty")
	ErrNegativeFileChange   = errors.New("[Config] Maximum file change must not be a negative number")
	ErrInvalidTitlePattern  = errors.New("[Config] Invalid pull request title pattern")
	ErrInvalidCommitPattern = errors.New("[Config] Invalid pull request commit message pattern")
	ErrInvalidBranchPattern = errors.New("[Config] Invalid pull request branch name pattern")
)

// Event error
var (
	ErrEventFileRead  = errors.New("[Event] Failed to read event file")
	ErrEventFileParse = errors.New("[Event] Failed to parse event file")
)

// validator error
var (
	ErrInvalidTitle   = errors.New("pull request title does not follow the desired pattern")
	ErrNoBody         = errors.New("pull request must have a non-empty body")
	ErrNoIssue        = errors.New("pull request does not mention any issues")
	ErrTooManyChanges = errors.New("pull request introduces too many changes")
	ErrInvalidBranch  = errors.New("pull request branch name does not follow the desired pattern")
)

// GitHub API error
var (
	ErrFirstComment = errors.New("no existing report was found")
)
