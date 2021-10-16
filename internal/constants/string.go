package constants

// Report template
var (
	ReportHeader      = `## Ethos' Run Report`
	WhitelistTemplate = `### Whitelist Report
	
%s`
	ValidatorTemplate = `### Validator Report
	
%s`
	WhitelistPass  = "Pull request matches with one (or more) enabled whitelist criteria. Pull request validation is skipped."
	WhitelistFail  = "Pull request does not match with all enabled whitelist criteria. Pull request will be validated."
	ValidationPass = "Pull request satisfies all enabled pull request rules."
	ValidationFail = "Pull request is invalid."
)

// Emojis

const (
	PassEmoji = "✅"
	FailEmoji = "❌"
)
