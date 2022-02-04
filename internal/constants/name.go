package constants

// Validator names
const (
	TitleValidatorName           = "Pull request has a valid title"
	IssueValidatorName           = "Pull request has mentioned issues"
	FileValidatorName            = "Pull request does not introduce too many changes"
	CommitValidatorName          = "All commits in this pull request has valid messages"
	VerifiedCommitsValidatorName = "All commits in this pull request come from verified users"
	BodyValidatorName            = "Pull request should have a non-empty body"
	BranchValidatorName          = "Pull request has valid branch name"
)

// Whitelist names
const (
	BotWhitelistName        = "Pull request is submitted by a bot and should be ignored"
	DraftWhitelistName      = "Pull request is a draft and should be ignored"
	PermissionWhitelistName = "Pull request is submitted by administrators and should be ignored"
	UsernameWhitelistName   = "Pull request is made by a whitelisted user and should be ignored"
)
