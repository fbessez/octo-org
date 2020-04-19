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
	if err != nil {
		forceRefresh = false
	}

	repoNames, err := fetchRepoNames(ctx, forceRefresh)
	spew.Dump(len(repoNames))
	check(err)

	orgStats  := make(models.OrgStats)
	var names [1]string
	names[0] = "guinness"
	for _, repoName := range names {
		stats, err := fetchRepoStats(ctx, forceRefresh, repoName)
		if err != nil {
			fmt.Println("error getting repo stats", repoName, err)
			continue
		}

		orgStats[repoName] = stats.Contributors
	}

	storeRepoStats(&orgStats)





// for i, repo := range repos {
	// get contribution stats
	// store contribution stats :::::: stats -> $repo_name -> $username -> $week
// }

// iterate through repos in stats
// collecting information on a per user basis
// adding together by week :::::: additions, deletions, commits

// allow stats to be sorted in every which way

	return
}

func main() {
	http.HandleFunc("/stats", statsHandler)
	fmt.Println("Listening on " + port)
	http.ListenAndServe(port, nil)
}