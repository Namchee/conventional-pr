package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/Namchee/ethos/internal"
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
	infoLogger = log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "[ERROR]", log.Ldate|log.Ltime)
}

func main() {
	ctx := context.Background()

	var config *entity.Config
	var meta *entity.Meta
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
	event, _ := entity.ReadEvent(
		utils.ReadEnvString("GITHUB_EVENT_PATH"),
	)

	pullRequest, err := client.GetPullRequest(ctx, meta.Owner, meta.Name, event.Number)

	if err != nil {
		errorLogger.Fatalln("Failed to fetch pull request data")
	}

	var vgResult []*entity.ValidationResult

	sync := &sync.WaitGroup{}

	infoLogger.Println("Testing pull request for whitelists")
	wg := whitelist.NewWhitelistGroup(client, config, meta, sync)
	wgResult := wg.Process(pullRequest)

	isWhitelisted := whitelist.IsWhitelisted(wgResult)

	if !isWhitelisted {
		infoLogger.Println("Testing pull request validity")
		vg := validator.NewValidatorGroup(client, config, meta, sync)
		vgResult = vg.Process(pullRequest)
	}

	svc := service.NewGithubService(client, config, meta)
	svc.WriteReport(pullRequest, wgResult, vgResult)
}
