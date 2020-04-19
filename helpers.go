package main

import (
	"context"
	"fmt"
	"net/http"
  "time"
  "encoding/json"
  "os"
  
  "go.opencensus.io/plugin/ochttp"
  // "github.com/davecgh/go-spew/spew"
  "github.com/fbessez/octo-org/github"
  "github.com/fbessez/octo-org/models"
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

func check(e error) {
  if e != nil {
     panic(e)
  }
}

func fetchRepoStats(ctx context.Context, forceRefresh bool, repoName string) (stats *models.GetContributerStatsByRepoResponse, err error) {
	if true {
		stats, err = githubClient.GetContributerStatsByRepo(ctx, repoName)
		check(err)
	} else {
		// Consult local storage
		// return
	}

	return stats, nil
}

func storeRepoStats(stats *models.OrgStats) (err error) {
	f, err := os.Create("./tmp/repo_name_to_stats.json")
	defer f.Close()
	check(err)

	bytes, err := json.Marshal(stats)
	n2, err := f.Write(bytes)
	fmt.Printf("wrote %d bytes", n2)

	return
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
	check(err)

	for _, repo := range response.Repos {
		repoNames = append(repoNames, repo.Name)
		err = rediscli.SetAdd(redisKeyRepoNames, repo.Name)
		if err != nil {
			fmt.Println("Setting " + repo.Name + "to " + redisKeyRepoNames + "failed miserably", err)
			continue
		}
	}

	return repoNames, nil
}
