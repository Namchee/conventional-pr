package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/Namchee/ethos/internal/service"
	"github.com/Namchee/ethos/internal/utils"
	"github.com/Namchee/ethos/internal/validator"
	"github.com/Namchee/ethos/internal/whitelist"
)

// Logger
var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lmsgprefix)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lmsgprefix)
}

func main() {
	infoLogger.Println("Initializing ethos")
	start := time.Now()

	ctx := context.Background()

	var config *entity.Config
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
		errorLogger.Fatalln(err)
	}

	infoLogger.Println("Initializing GitHub Client")
	client := internal.NewGithubClient(config)

	infoLogger.Println("Reading pull request metadata")
	event, err = entity.ReadEvent(os.DirFS(""))

	if err != nil {
		errorLogger.Fatalln(err)
	}

	infoLogger.Println("Validating pull request sub-events")
	if !utils.ContainsString(constants.Events, event.Action) {
		infoLogger.Println("Incompatible sub-events detected. Exiting...")
		os.Exit(0)
	}

	pullRequest, err := client.GetPullRequest(ctx, meta.Owner, meta.Name, event.Number)

	if err != nil {
		errorLogger.Fatalln("Failed to fetch pull request data")
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
	err = svc.WriteReport(pullRequest, wgResult, vgResult)

	if err != nil {
		errorLogger.Fatalln("Failed to write report")
	}

	if !validator.IsValid(vgResult) {
		infoLogger.Println("Processing invalid pull request")

		err = svc.AttachLabel(pullRequest)
		if err != nil {
			errorLogger.Fatalln("Failed to attach invalid pull request label")
		}

		err = svc.WriteTemplate(pullRequest)
		if err != nil {
			errorLogger.Fatalln("Failed to write message template")
		}

		err = svc.ClosePullRequest(pullRequest)
		if err != nil {
			errorLogger.Fatalln("Failed to close invalid pull request")
		}

		infoLogger.Printf("Finished processing on %.2fs", time.Since(start).Seconds())
		os.Exit(1)
	}

	infoLogger.Printf("Finished processing on %.2fs", time.Since(start).Seconds())
}
