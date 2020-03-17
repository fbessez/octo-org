package github

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *GithubClientImpl) formRequest(action string, url string) (request *http.Response, err error) {
	request, err := http.NewRequest("GET", getOrgUrl())
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(config.Constants.Username, config.Constants.Password)

	return
}

func (c *GithubClientImpl) makeRequest(request *http.Response) (request *http.Response, err error) {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
}


func (c *GithubClientImpl) unmarshalResponse(ctx context.Context, resp *http.Response, i interface{}) (err error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		return err
	}

	return
}
