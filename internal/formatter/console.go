package formatter

import (
	"fmt"
	"strings"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func formatWhitelistResultToConsole(
	whitelistResults []*entity.WhitelistResult,
) string {

	t := table.NewWriter()

	t.SetStyle(table.StyleLight)
	t.SetTitle("Whitelist Result")
	t.AppendHeader(table.Row{"Whitelist", "Active", "Result"})

	flag := false

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

		t.AppendRow(table.Row{r.Name, active, verdict})
	}

	summary := constants.WhitelistFail
	if flag {
		summary = constants.WhitelistPass
	}

	return fmt.Sprintf(
		"%s\n\nResult: %s",
		t.Render(),
		summary,
	)
}

func formatValidationResultToConsole(
	validationResults []*entity.ValidationResult,
) string {
	var errors []error

	t := table.NewWriter()

	t.SetStyle(table.StyleLight)
	t.SetTitle("Validation Result")
	t.AppendHeader(table.Row{"Validation", "Active", "Result"})

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

		t.AppendRow(table.Row{r.Name, active, verdict})
	}

	var reasons []string
	verdict := constants.ValidationPass

	if len(errors) > 0 {
		verdict = constants.ValidationFail

		for _, fail := range errors {
			reasons = append(reasons, fmt.Sprintf("- %s", utils.Capitalize(fail.Error())))
		}
	}

	report := fmt.Sprintf(
		"%s\n\nResult: %s",
		t.Render(),
		verdict,
	)

	if len(reasons) > 0 {
		report = fmt.Sprintf(
			"%s\n\nReason:\n%s",
			report,
			strings.Join(reasons, "\n"),
		)
	}

	return report
}

// FormatResultToTables formats both whitelist and validation results for workflow reporting to console
func FormatResultToConsole(
	results *entity.PullRequestResult,
) string {
	report := fmt.Sprintf(
		"%s\n\n%s",
		constants.LogHeader,
		formatWhitelistResultToConsole(results.Whitelist),
	)

	if len(results.Validation) > 0 {
		report = fmt.Sprintf(
			"%s\n\n%s",
			report,
			formatValidationResultToConsole(results.Validation),
		)
	}

	return report
}
