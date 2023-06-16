package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/utils"
)

func formatWhitelistResultToTable(
	whitelistResults []*entity.WhitelistResult,
) string {
	header := "| Whitelist | Active | Result |"
	separator := "| - | :-: | :-: |"

	flag := false
	results := []string{
		header,
		separator,
	}

	for _, r := range whitelistResults {
		active := constants.FailEmoji
		verdict := constants.FailEmoji

		if r.Active {
			active = constants.PassEmoji
		}

		if r.Result {
			flag = true
			verdict = constants.PassEmoji
		}

		result := fmt.Sprintf("| %s | %s | %s |", r.Name, active, verdict)
		results = append(results, result)
	}

	summary := constants.WhitelistFail
	if flag {
		summary = constants.WhitelistPass
	}

	report := fmt.Sprintf(constants.ResultTemplate, summary)

	return fmt.Sprintf(
		constants.WhitelistTemplate,
		strings.Join(results, "\n"),
		report,
	)
}

func formatValidationResultToTable(
	validationResults []*entity.ValidationResult,
) string {
	header := "| Validation | Active | Result |"
	separator := "| - | :-: | :-: |"

	var errors []error
	results := []string{
		header,
		separator,
	}

	for _, r := range validationResults {
		active := constants.PassEmoji
		verdict := constants.PassEmoji

		if !r.Active {
			active = constants.FailEmoji
		}

		if r.Result != nil {
			errors = append(errors, r.Result)
			verdict = constants.FailEmoji
		}

		result := fmt.Sprintf("| %s | %s | %s |", r.Name, active, verdict)
		results = append(results, result)
	}

	var reason string
	verdict := constants.ValidationPass

	if len(errors) > 0 {
		verdict = constants.ValidationFail
		var reasons []string

		for _, fail := range errors {
			reasons = append(reasons, fmt.Sprintf("- %s", utils.Capitalize(fail.Error())))
		}
		reason = fmt.Sprintf(constants.ReasonTemplate, strings.Join(reasons, "\n"))
	}

	report := fmt.Sprintf(constants.ResultTemplate, verdict)

	result := fmt.Sprintf(constants.ValidatorTemplate, strings.Join(results, "\n"), report)

	if reason != "" {
		result = fmt.Sprintf("%s\n\n%s", result, reason)
	}

	return result
}

// FormatResultToTables formats both whitelist and validation results for workflow reporting in markdown syntax
func FormatResultToTables(
	results *entity.PullRequestResult,
	now time.Time,
) string {
	report := fmt.Sprintf(
		"%s\n\n%s",
		constants.ReportHeader,
		formatWhitelistResultToTable(
			results.Whitelist,
		),
	)

	if len(results.Validation) > 0 {
		report = fmt.Sprintf(
			"%s\n\n%s",
			report,
			formatValidationResultToTable(
				results.Validation,
			),
		)
	}

	return fmt.Sprintf(
		"%s\n\n<sub>Last Modified at %s</sub>",
		report,
		now.Format(time.RFC822),
	)
}
