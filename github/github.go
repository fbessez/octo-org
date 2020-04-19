package github

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/fbessez/octo-org/models"
	"github.com/fbessez/octo-org/config"
)

const (
	GITHUB_BASE_URL = "https://api.github.com/"
)

type GithubClient struct {
	HttpClient *http.Client
}

func (c *GithubClient) GetUser(ctx context.Context) (getUserResponse *models.GetUserResponse, err error) {
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

	getUserResponse = &models.GetUserResponse{}
	c.unmarshalResponse(ctx, resp, getUserResponse)
	spew.Dump(getUserResponse)

	return
}

func (c *GithubClient) GetOrg(ctx context.Context) (getOrgResponse *models.GetOrgResponse, err error) {
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

	getOrgResponse = &models.GetOrgResponse{}
	c.unmarshalResponse(ctx, resp, getOrgResponse)
	spew.Dump(getOrgResponse)

	return
}

func (c *GithubClient) GetReposByOrgByPage(ctx context.Context, page int) (getReposByOrgResponse *models.GetReposByOrgResponse, err error) {
	url := GITHUB_BASE_URL + "orgs/" + config.CONSTANTS.OrgName + "/repos?per_page=100&page=" + strconv.Itoa(page)

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
	c.unmarshalResponse(ctx, resp, &getReposByOrgResponse.Repos)

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
	c.unmarshalResponse(ctx, resp, &getContributerStatsByRepoResponse.Contributors)

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

func (c *GithubClient) GetAllReposByOrg(ctx context.Context) (getReposByOrgResponse *models.GetReposByOrgResponse, err error) {
	getReposByOrgResponse = &models.GetReposByOrgResponse{}

	for page := 1; page < 3; page++ {
		r, err := c.GetReposByOrgByPage(ctx, page)
		if err != nil {
			fmt.Println("error paginating", err)
			return nil, err
		}

		getReposByOrgResponse.Repos = append(getReposByOrgResponse.Repos, r.Repos...)
	}

	return
}