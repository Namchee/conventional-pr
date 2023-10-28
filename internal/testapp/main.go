package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/entity"
)

func main() {
	token := ""

	client := internal.NewGithubClient(&entity.Configuration{Token: token, GraphQLURL: "https://api.github.com/graphql"})
	items, err := client.GetCommits(context.Background(), &entity.Meta{
		Owner: "Namchee",
		Name:  "conventional-pr",
	}, 77)

	fmt.Println(err)

	beautified, _ := json.MarshalIndent(items, "", "  ")
	fmt.Println(string(beautified))
}
