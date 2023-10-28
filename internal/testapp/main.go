package main

import (
	"context"
	"fmt"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/entity"
)

func main() {
	token := "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"

	client := internal.NewGithubClient(&entity.Configuration{Token: token, GraphQLURL: "https://api.github.com/graphql"})
	_, err := client.GetPullRequest(context.Background(), &entity.Meta{
		Owner: "Namchee",
		Name:  "conventional-pr",
	}, 95)

	fmt.Println(err)
}
