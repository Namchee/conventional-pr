package constants

// Validator names
const (
	TitleValidatorName           = "Pull request has a valid title"
	IssueValidatorName           = "Pull request mentioned issue(s)"
	FileValidatorName            = "Pull request does not introduce too much changes"
	CommitValidatorName          = "All commit(s) in this pull request has valid messages"
	VerifiedCommitsValidatorName = "All commit(s) in this pull request come from verified user(s)"
	BodyValidatorName            = "Pull request should have a non-empty body"
	BranchValidatorName          = "Pull request has valid branch name"
)

// Whitelist names
const (
	BotWhitelistName        = "Pull request is submitted by a bot and should be ignored"
	DraftWhitelistName      = "Pull request is still a draft and should be ignored"
	PermissionWhitelistName = "Pull request is submitted by user with high privileges and should be ignored"
)
