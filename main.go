package main
	
import (
	"fmt"
  "net/http"
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

	repoNames, err := fetchRepoNames(ctx, forceRefresh)
	check(err)

	orgStats, err := getOrgStats(ctx, forceRefresh, repoNames)
	check(err)

	orgStatsByUser, err := models.ConvertOrgStatsToOrgStatsByUser(*orgStats)
	check(err)

	writeUserStats(orgStatsByUser)

// allow stats to be sorted in every which way

	return
}

func main() {
	spew.Dump(port)
	http.HandleFunc("/stats", statsHandler)
	fmt.Println("Listening on " + port)
	http.ListenAndServe(port, nil)
}