package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/fbessez/orgContributions/models"
	"github.com/fbessez/orgContributions/config"
)

const (
	GITHUB_BASE_URL = "https://api.github.com/"
)

type GithubClient struct {
	HttpClient *http.Client
}

func (c *GithubClient) GetUser(ctx context.Context) (getUserResponse *models.GetUserResponse, err error) {
	url := GITHUB_BASE_URL + "orgs/" + config.CONSTANTS.OrgName

	req,  err := c.formRequest("GET", url)
	if err != nil { 
		fmt.Println("error forming request", err)
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		fmt.Println("error making request", err)
		return nil, err
	}

	getUserResponse = &models.GetUserResponse{}
	c.unmarshalResponse(ctx, resp, getUserResponse)
	spew.Dump(getUserResponse)

	return
}

func (c *GithubClient) GetOrg(ctx context.Context) (getOrgResponse *models.GetOrgResponse, err error) {
	url := GITHUB_BASE_URL + "user"

	req,  err := c.formRequest("GET", url)
	if err != nil { 
		fmt.Println("error forming request", err)
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		fmt.Println("error making request", err)
		return nil, err
	}

	getOrgResponse = &models.GetOrgResponse{}
	c.unmarshalResponse(ctx, resp, getOrgResponse)
	spew.Dump(getOrgResponse)

	return
}

func (c *GithubClient) GetReposByOrg(ctx context.Context) (getReposByOrgResponse *models.GetReposByOrgResponse, err error) {
	url := GITHUB_BASE_URL + "orgs/" + config.CONSTANTS.OrgName + "/repos"

	req,  err := c.formRequest("GET", url)
	if err != nil { 
		fmt.Println("error forming request", err)
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		fmt.Println("error making request", err)
		return nil, err
	}

	getReposByOrgResponse = &models.GetReposByOrgResponse{}
	c.unmarshalResponse(ctx, resp, getReposByOrgResponse)
	spew.Dump(getReposByOrgResponse)

	return
}

func (c *GithubClient) GetContributerStatsByRepo(ctx context.Context, repoName string) (getContributerStatsByRepoResponse *models.GetContributerStatsByRepoResponse, err error) {
	url := GITHUB_BASE_URL + "repos/" + config.CONSTANTS.OrgName + "/" + repoName + "/stats/contributors"

	req,  err := c.formRequest("GET", url)
	if err != nil { 
		fmt.Println("error forming request", err)
		return nil, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		fmt.Println("error making request", err)
		return nil, err
	}

	getContributerStatsByRepoResponse = &models.GetContributerStatsByRepoResponse{}
	c.unmarshalResponse(ctx, resp, getContributerStatsByRepoResponse)
	spew.Dump(getContributerStatsByRepoResponse)

	return
}

func (c *GithubClient) CheckOrgMembership(ctx context.Context, username string) (isMember bool, err error) {
	url := GITHUB_BASE_URL + "orgs/" + config.CONSTANTS.OrgName + "/members/" + username

	req,  err := c.formRequest("GET", url)
	if err != nil { 
		fmt.Println("error forming request", err)
		return false, err
	}

	resp, err := c.makeRequest(req)
	if err != nil {
		fmt.Println("error making request", err)
		return false, err
	}

	spew.Dump(resp)
	return resp.StatusCode == 200, nil
}

