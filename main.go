package main
	
import (
	"encoding/json"
	"fmt"
  "net/http"
  "sort"

  "github.com/davecgh/go-spew/spew"
  "github.com/fbessez/octo-org/models"
)

const (
	port = ":8090"
)

var (
	availableSorting = [4]string {"additions", "deletions", "commits", "ratio"}
)

func statsHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	enableCors(&w)

	forceRefresh := getForceRefreshValue(req)

	sortOption := req.URL.Query().Get("sort")
	if sortOption == "" { sortOption = "commits" }

	repoNames, err := getRepoNames(ctx, forceRefresh)
	check(err)

	orgStats, err := getOrgStats(ctx, forceRefresh, repoNames)
	check(err)

	orgStatsByUser, err := models.ConvertOrgStatsToOrgStatsByUser(*orgStats)
	check(err)

	writeUserStats(orgStatsByUser)

	userCommits := getUserCommits(*orgStatsByUser)

	switch sortOption {
	case "additions":
		sort.Slice(userCommits, func(i, j int) bool {
			return userCommits[i].TotalAdditions > userCommits[j].TotalAdditions
		})
	case "deletions": 
		sort.Slice(userCommits, func(i, j int) bool {
			return userCommits[i].TotalDeletions > userCommits[j].TotalDeletions
		})
	default: 
		sort.Slice(userCommits, func(i, j int) bool {
			return userCommits[i].TotalCommits > userCommits[j].TotalCommits
		})
	}

	json.NewEncoder(w).Encode(userCommits)

	return
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	spew.Dump(port)
	http.HandleFunc("/stats", statsHandler)
	fmt.Println("Listening on " + port)
	http.ListenAndServe(port, nil)
}