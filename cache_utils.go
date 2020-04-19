package main

import (
	"fmt"
  "encoding/json"
  "io/ioutil"
  "os"
 
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
			fmt.Println("Setting " + repoName + "to " + redisKeyRepoNames + "failed miserably", err)
			continue
		}
	}
}

func readRepoNames() (repoNames []string, err error) {
	repoNames, err = rediscli.GetSetMembers(redisKeyRepoNames)
	check(err)

	return
}




