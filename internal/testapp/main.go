package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/entity"
)

func main() {
	url, _ := url.Parse("https://api.github.com/graphql")
	token := ""

	client := internal.NewGithubClient(&entity.Configuration{ Token: token, BaseURL: url })
	_, err := client.GetPullRequest(context.Background(), &entity.Meta{
		Owner: "Namchee",
		Name: "conventional-pr",
	}, 95)

	fmt.Println(err)
}
