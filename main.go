package main
	
import (
	"fmt"
	"time"
  "net/http"
  "go.opencensus.io/plugin/ochttp"

  "github.com/davecgh/go-spew/spew"
  "github.com/fbessez/orgContributions/github"
)

const (
	port = ":8090"
)

func newGithubClient() *github.GithubClient {
	var httpClient = &http.Client{Transport: &ochttp.Transport{}, Timeout: 5 * time.Second}
	return &github.GithubClient{HttpClient: httpClient}
}

func getOrgStats(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	githubClient := newGithubClient()

	resp, err := githubClient.GetOrg(ctx)
	if err != nil {
		fmt.Println("error getting Org", err)
		return
	}

	return
}

func main() {
	http.HandleFunc("/getOrgStats", getOrgStats)
	fmt.Println("Listening on " + port)
	http.ListenAndServe(port, nil)
}