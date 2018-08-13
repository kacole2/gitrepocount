package function

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

//Struct to Marshall JSON response
type RepoCount struct {
	User  string `json:"user"`
	Count int    `json:"repocount"`
}

//Standard OpenFaaS func to get Secrets from K8s
func getAPISecret(secretName string) (secretBytes []byte, err error) {
	// read from the openfaas secrets folder
	secretBytes, err = ioutil.ReadFile("/var/openfaas/secrets/" + secretName)
	if err != nil {
		// read from the original location for backwards compatibility with openfaas <= 0.8.2
		secretBytes, err = ioutil.ReadFile("/run/secrets/" + secretName)
	}

	return secretBytes, err
}

// Handle a serverless request
func Handle(req []byte) string {
	input := string(req)

	//Get the GitHub API Access Key from Kubernetes Secrets
	secretBytes, err := getAPISecret("github-api-secret")
	if err != nil {
		log.Fatal(err)
	}
	secret := strings.TrimSpace(string(secretBytes))

	//Use the Secret Token to Create a New *Authenticated* Client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: secret},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	//Pagination Option for every 20 records
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 20},
	}

	var allRepos []*github.Repository
	//Loop that will continue counting repos every 20 until complete
	for {
		repos, resp, err := client.Repositories.List(ctx, string(input), opt)
		if err != nil {
			log.Fatalf("User %s doesn't exist. Error: %s", string(req), err.Error())
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	//Marshal response into JSON
	totalCount := &RepoCount{User: string(req), Count: len(allRepos)}
	repoCountMarsh, err := json.Marshal(totalCount)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return fmt.Sprintf("%s", string(repoCountMarsh))
}
