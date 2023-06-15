package formatter

import (
	"fmt"
	"log"
	"os"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func formatWhitelistResultToConsole(
	whitelistResults []*entity.WhitelistResult,
	logger *log.Logger,
) {

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
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

	t.Render()

	logger.Println()
	logger.Printf("Result: %s\n", summary)
}

func formatValidationResultToConsole(
	validationResults []*entity.ValidationResult,
	logger *log.Logger,
) {
	var errors []error

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
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

	t.Render()

	logger.Println()
	logger.Printf("Result: %s\n", verdict)

	for _, reason := range reasons {
		logger.Println(reason)
	}
}

// FormatResultToTables formats both whitelist and validation results for workflow reporting to console
func FormatResultToConsole(
	whitelistResults []*entity.WhitelistResult,
	validationResults []*entity.ValidationResult,
	logger *log.Logger,
) {
	logger.Println(constants.LogHeader)
	logger.Println("------------------------------")
	logger.Println(constants.LogSubtitle)
	logger.Println()

	formatWhitelistResultToConsole(whitelistResults, logger)

	if len(validationResults) > 0 {
		logger.Println()
		formatValidationResultToConsole(validationResults, logger)
	}
}
