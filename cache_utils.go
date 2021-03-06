package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/fbessez/octo-org/models"
	"github.com/fbessez/octo-org/rediscli"
)

const (
	repo_stats_file = "./tmp/repo_name_to_stats.json"
	user_stats_file = "./tmp/user_name_to_stats.json"
)

func writeRepoStats(orgStats *models.OrgStats) (err error) {
	f, err := os.Create(repo_stats_file)
	defer f.Close()
	check(err)

	bytes, err := json.Marshal(orgStats)
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

func writeUserStats(userStats *models.OrgStatsByUser) (err error) {
	f, err := os.Create(user_stats_file)
	defer f.Close()
	check(err)

	bytes, err := json.Marshal(userStats)
	n2, err := f.Write(bytes)
	fmt.Printf("wrote %d bytes", n2)

	return
}

func readUserStats() (userStats *models.OrgStatsByUser, err error) {
	jsonFile, err := os.Open(user_stats_file)
	defer jsonFile.Close()
	check(err)

	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &userStats)

	return userStats, nil
}

func writeRepoNames(repoNames []string) {
	for _, repoName := range repoNames {
		err := rediscli.SetAdd(redisKeyRepoNames, repoName)
		if err != nil {
			fmt.Println("Setting "+repoName+"to "+redisKeyRepoNames+"failed miserably", err)
			continue
		}
	}
}

func readRepoNames() (repoNames []string, err error) {
	repoNames, err = rediscli.GetSetMembers(redisKeyRepoNames)
	check(err)

	return
}

func getForceRefreshValue(req *http.Request) bool {
	forceRefresh, err := strconv.ParseBool(req.URL.Query().Get("forceRefresh"))
	if forceRefresh != true || err != nil {
		if !fileExists(repo_stats_file) || !fileExists(user_stats_file) {
			forceRefresh = true
		} else {
			forceRefresh = false
		}
	}

	return forceRefresh
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
