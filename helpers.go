package main

import (
	"context"
	"fmt"
	"net/http"
	// "sort"
  "time"
  
  "go.opencensus.io/plugin/ochttp"

  "github.com/fbessez/octo-org/github"
  "github.com/fbessez/octo-org/models"
  "github.com/fbessez/octo-org/config"
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

func getRepoNames(ctx context.Context, forceRefresh bool) (repoNames []string, err error) {
	if forceRefresh {
		repoNames, err = refreshAllRepoNames(ctx)
		check(err)

		writeRepoNames(repoNames)
		return repoNames, nil
	}

	repoNames, err = readRepoNames()
	check(err)

	return repoNames, nil
}

func refreshAllRepoNames(ctx context.Context) (repoNames []string, err error) {
	repos, err := fetchRepoNames(ctx)

	for _, repo := range repos {
		repoNames = append(repoNames, repo.Name)
	}

	return repoNames, nil
}

func fetchRepoNames(ctx context.Context) (repos []*models.Repository, err error) {
	response, err := githubClient.GetAllReposByOrg(ctx)
	check(err)

	return response.Repos, nil
}

func getOrgStats(ctx context.Context, forceRefresh bool, repoNames []string) (orgStats *models.OrgStats, err error) {
	if forceRefresh {
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

	for _, repoName := range repoNames {
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

func getUserCommits(orgStatsByUser models.OrgStatsByUser) (userCommits []*models.UserCommits) {
	for githubUsername, userStats := range orgStatsByUser {
		totalCommits := 0
		for _, repoStats := range *userStats {
			totalCommits += repoStats.TotalCommits
		}

		userCommits = append(userCommits, &models.UserCommits{GithubUsername: githubUsername, TotalCommits: totalCommits})
	}
	
	return
}

