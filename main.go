package main
	
import (
	"encoding/json"
	"fmt"
  "net/http"
  "sort"
  "strconv"

  "github.com/fbessez/octo-org/models"
  "github.com/davecgh/go-spew/spew"
)

const (
	port = ":8090"
)

func statsHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	forceRefresh, err := strconv.ParseBool(req.URL.Query().Get("forceRefresh"))
	if err != nil { forceRefresh = false }

	repoNames, err := getRepoNames(ctx, forceRefresh)
	check(err)

	orgStats, err := getOrgStats(ctx, forceRefresh, repoNames)
	check(err)

	orgStatsByUser, err := models.ConvertOrgStatsToOrgStatsByUser(*orgStats)
	check(err)

	writeUserStats(orgStatsByUser)

	userCommits := getUserCommits(*orgStatsByUser)
	sort.Slice(userCommits, func(i, j int) bool {
		return userCommits[i].TotalCommits > userCommits[j].TotalCommits
	})

	json.NewEncoder(w).Encode(userCommits)

	return
}

func main() {
	spew.Dump(port)
	http.HandleFunc("/stats", statsHandler)
	fmt.Println("Listening on " + port)
	http.ListenAndServe(port, nil)
}