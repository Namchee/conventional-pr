package formatter

import (
	"fmt"
	"log"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/utils"
)

func formatWhitelistResultToConsole(
	whitelistResults []*entity.WhitelistResult,
	logger *log.Logger,
) {
	flag := false

	for _, r := range whitelistResults {
		active := constants.InactiveLabel
		verdict := constants.FailLabel

		if r.Active {
			active = constants.ActiveLabel
		}

		if r.Result {
			flag = true
			verdict = constants.PassLabel
		}

		logger.Printf("%s — %s — %s\n", r.Name, active, verdict)
	}

	summary := constants.WhitelistFail
	if flag {
		summary = constants.WhitelistPass
	}

	logger.Println(summary)
}

func formatValidationResultToConsole(
	validationResults []*entity.ValidationResult,
	logger *log.Logger,
) {
	var errors []error

	for _, r := range validationResults {
		active := constants.ActiveLabel
		verdict := constants.PassLabel

		if !r.Active {
			active = constants.InactiveLabel
		}

		if r.Result != nil {
			errors = append(errors, r.Result)
			verdict = constants.FailLabel
		}

		logger.Printf("%s — %s — %s\n", r.Name, active, verdict)
	}

	var reasons []string
	verdict := constants.ValidationPass

	if len(errors) > 0 {
		verdict = constants.ValidationFail

		for _, fail := range errors {
			reasons = append(reasons, fmt.Sprintf("- %s", utils.Capitalize(fail.Error())))
		}
	}

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

	formatWhitelistResultToConsole(whitelistResults, logger)

	if len(validationResults) > 0 {
		formatValidationResultToConsole(validationResults, logger)
	}
}
