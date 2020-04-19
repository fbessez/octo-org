package main

import (
	"context"
	"fmt"
	"net/http"
  "time"
  "encoding/json"
  "io/ioutil"
  "os"
  
  "go.opencensus.io/plugin/ochttp"
  "github.com/davecgh/go-spew/spew"
  "github.com/fbessez/octo-org/github"
  "github.com/fbessez/octo-org/models"
  "github.com/fbessez/octo-org/config"
  "github.com/fbessez/octo-org/rediscli"
  "github.com/gomodule/redigo/redis"
)

const (
	repo_stats_file = "./tmp/repo_name_to_stats.json"
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

func getOrgStats(ctx context.Context, forceRefresh bool, repoNames []string) (orgStats *models.OrgStats, err error) {
	if false {
		orgStats, err := refreshAllRepoStats(ctx, forceRefresh, repoNames)
		check(err)
		writeRepoStats(orgStats)

		return orgStats, nil
	}

	orgStats, err = readRepoStats()
	check(err)

	return orgStats, nil
}

func refreshAllRepoStats(ctx context.Context, forceRefresh bool, repoNames []string) (orgStats *models.OrgStats, err error) {
	result := make(models.OrgStats)
	spew.Dump(repoNames[0])

	var names [1]string
	names[0] = "guinness"
	for _, repoName := range names {
		stats, err := fetchRepoStats(ctx, repoName)
		if err != nil {
			fmt.Println("error getting repo stats", repoName, err)
			continue
		}

		result[repoName] = stats.Contributors
	}

	return &result, nil
}

func fetchRepoStats(ctx context.Context, repoName string) (stats *models.GetContributerStatsByRepoResponse, err error) {
	stats, err = githubClient.GetContributerStatsByRepo(ctx, repoName)
	check(err)

	return stats, nil
}

func writeRepoStats(stats *models.OrgStats) (err error) {
	f, err := os.Create(repo_stats_file)
	defer f.Close()
	check(err)

	bytes, err := json.Marshal(stats)
	n2, err := f.Write(bytes)
	fmt.Printf("wrote %d bytes", n2)

	return
}

func readRepoStats() (orgStats *models.OrgStats, err error) {
	jsonFile, err := os.Open(repo_stats_file)
	defer jsonFile.Close()
	check(err)

	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &orgStats)

	return orgStats, nil
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
