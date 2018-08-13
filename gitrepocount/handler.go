package function

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/go-github/github"
)

type RepoCount struct {
	User  string `json:"user"`
	Count int    `json:"repocount"`
}

// Handle a serverless request
func Handle(req []byte) string {
	input := string(req)
	client := github.NewClient(nil)

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 20},
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(context.Background(), string(input), opt)
		if err != nil {
			log.Fatalf("User %s doesn't exist. Error: %s", string(req), err.Error())
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	totalCount := &RepoCount{User: string(req), Count: len(allRepos)}
	repoCountMarsh, err := json.Marshal(totalCount)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return fmt.Sprintf("%s", string(repoCountMarsh))
}
