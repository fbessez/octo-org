package main

import (
	"context"
	"fmt"
	"net/http"
  "time"
  
  "go.opencensus.io/plugin/ochttp"

  "github.com/davecgh/go-spew/spew"
  "github.com/fbessez/octo-org/github"
  "github.com/fbessez/octo-org/config"
  "github.com/fbessez/octo-org/rediscli"
  "github.com/gomodule/redigo/redis"
)

var redisKeyRepoNames = config.CONSTANTS.OrgName + "::repos"
var githubClient = newGithubClient()

func newGithubClient() *github.GithubClient {
	var httpClient = &http.Client{Transport: &ochttp.Transport{}, Timeout: 5 * time.Second}
	return &github.GithubClient{HttpClient: httpClient}
}

func fetchRepoNames(ctx context.Context, forceRefresh bool) (repoNames []string, err error) {
	if forceRefresh {
		repoNames, err = getAndSetRepoNames(ctx)
		if err != nil {
			fmt.Println("Getting repositories by org failed miserably", err)
			return nil, err
		}
	} else {
		repoNames, err = rediscli.GetSetMembers(redisKeyRepoNames)
		if err != nil && err == redis.ErrNil {
			repoNames, err = getAndSetRepoNames(ctx)
			if err != nil {
				fmt.Println("Getting repositories by org failed miserably", err)
				return nil, err
			}
		}
	}

	return repoNames, nil
}

func getAndSetRepoNames(ctx context.Context) (repoNames []string, err error) {
	response, err := githubClient.GetAllReposByOrg(ctx)
	if err != nil {
		fmt.Println("Getting repositories by org failed miserably", err, response)
		return nil, err
	}

	for i, repo := range response.Repos {
		spew.Dump(i)
		repoNames = append(repoNames, repo.Name)
		err = rediscli.SetAdd(redisKeyRepoNames, repo.Name)
		if err != nil {
			fmt.Println("Setting " + repo.Name + "to " + redisKeyRepoNames + "failed miserably", err)
			continue
		}
	}

	return repoNames, nil
}