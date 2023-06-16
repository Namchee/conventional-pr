package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/Namchee/conventional-pr/internal/formatter"
	"github.com/Namchee/conventional-pr/internal/service"
	"github.com/Namchee/conventional-pr/internal/utils"
	"github.com/Namchee/conventional-pr/internal/validator"
	"github.com/Namchee/conventional-pr/internal/whitelist"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger

	defaultLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lmsgprefix)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lmsgprefix)
	defaultLogger = log.New(os.Stdout, "", 0)
}

func main() {
	infoLogger.Println("Initializing conventional-pr")
	start := time.Now()

	ctx := context.Background()

	var config *entity.Configuration
	var meta *entity.Meta
	var event *entity.Event
	var err error

	infoLogger.Println("Reading configuration from environment variables")
	config, err = entity.ReadConfig()

	if err != nil {
		errorLogger.Fatalln(err)
	}

	infoLogger.Println("Reading repository metadata")
	meta, err = entity.CreateMeta(
		utils.ReadEnvString("GITHUB_REPOSITORY"),
	)

	if err != nil {
		errorLogger.Fatalf("Failed to read repository metadata: %s", err.Error())
	}

	infoLogger.Println("Initializing GitHub Client")
	client := internal.NewGithubClient(config)

	infoLogger.Println("Reading pull request metadata")
	event, err = entity.ReadEvent(os.DirFS("/"))

	if err != nil {
		errorLogger.Fatalf("Failed to read repository event: %s", err.Error())
	}

	infoLogger.Println("Validating pull request sub-events")
	if !utils.ContainsString(constants.Events, event.Action) {
		infoLogger.Println("Incompatible sub-events detected. Exiting...")
		os.Exit(0)
	}

	pullRequest, err := client.GetPullRequest(ctx, meta.Owner, meta.Name, event.Number)

	if err != nil {
		errorLogger.Fatalf("Failed to fetch pull request data: %s", err.Error())
	}

	var vgResult []*entity.ValidationResult

	sync := &sync.WaitGroup{}

	infoLogger.Println("Testing pull request for whitelists")
	wg := whitelist.NewWhitelistGroup(client, config, meta, sync)
	wgResult := wg.Process(pullRequest)

	if !whitelist.IsWhitelisted(wgResult) {
		infoLogger.Println("Testing pull request validity")
		vg := validator.NewValidatorGroup(client, config, meta, sync)
		vgResult = vg.Process(pullRequest)
	}

	svc := service.NewGithubService(client, config, meta)

	infoLogger.Println("Writing run report")

	results := &entity.PullRequestResult{
		Whitelist:  wgResult,
		Validation: vgResult,
	}

	resultLog := formatter.FormatResultToConsole(results)
	defaultLogger.Println(resultLog)

	if config.Verbose {
		err = svc.WriteReport(pullRequest, results, time.Now())

		if err != nil {
			errorLogger.Fatalf("Failed to write report: %s", err.Error())
		}
	}

	if !validator.IsValid(vgResult) {
		infoLogger.Println("Processing invalid pull request")

		err = svc.AttachLabel(pullRequest)
		if err != nil {
			errorLogger.Fatalf("Failed to attach invalid pull request label: %s", err.Error())
		}

		err = svc.WriteMessage(pullRequest)
		if err != nil {
			errorLogger.Fatalf("Failed to write message: %s", err.Error())
		}

		err = svc.ClosePullRequest(pullRequest)
		if err != nil {
			errorLogger.Fatalf("Failed to close invalid pull request: %s", err.Error())
		}

		infoLogger.Printf("Finished processing on %.2fs", time.Since(start).Seconds())
		os.Exit(1)
	}

	infoLogger.Printf("Finished processing on %.2fs", time.Since(start).Seconds())
}
