package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"


	"github.com/davecgh/go-spew/spew"
	"github.com/fbessez/orgContributions/config"
)

func (c *GithubClient) formRequest(action string, url string) (request *http.Request, err error) {
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(config.CONSTANTS.Username, config.CONSTANTS.Password)

	return
}

func (c *GithubClient) makeRequest(request *http.Request) (response *http.Response, err error) {
	response, err = c.HttpClient.Do(request)
	if err != nil {
		fmt.Println("error making http request", err)
		return nil, err
	}

	spew.Dump(response.Status)

	return
}


func (c *GithubClient) unmarshalResponse(ctx context.Context, resp *http.Response, i interface{}) (err error) {
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("error closing response body", err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body", err)
		return err
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		fmt.Println("error unmarshalling json", err)
		return err
	}

	return nil
}
