package github

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"net/http"

	"github.com/fbessez/orgContributions/models"
)

const (
	GITHUB_BASE_URL = "https://api.github.com/"
)

type GithubClientImpl struct {
	httpClient *http.Client
}

func (c *GithubClientImpl) GetUser(ctx context.Context) (getUserResponse *models.GetUserResponse, err error) {
	url := GITHUB_BASE_URL + "orgs/" + config.Constants.OrgName
	request,  err := c.formRequest("GET", url)
	response, err := c.makeRequest(request)

	c.unmarshalResponse(ctx, resp, getUserResponse)

	return
}

func (c *GithubClientImpl) GetOrg(ctx context.Context) (getOrgResponse *models.GetOrgResponse, err error) {
	url := GITHUB_BASE_URL + "user"
	req,  err := c.formRequest("GET", url)
	resp, err := c.makeRequest(req)

	c.unmarshalResponse(ctx, resp, getOrgResponse)

	return
}

func (c *GithubClientImpl) GetReposByOrg(ctx context.Context) (getReposByOrgResponse *models.GetReposByOrgResponse, err error) {
	url := GITHUB_BASE_URL + "orgs/" + config.Constants.OrgName + "/repos"
	req,  err := c.formRequest("GET", url)
	resp, err := c.makeRequest(req)

	c.unmarshalResponse(ctx, resp, getReposByOrgResponse)

	return
}

func (c *GithubClientImpl) GetContributerStatsByRepo(ctx context.Context, repoName string) (getContributerStatsByRepoResponse *models.GetContributerStatsByRepoResponse, err error) {
	url := GITHUB_BASE_URL + "repos/" + config.Constants.OrgName + "/" + repoName + "/stats/contributors"
	req,  err := c.formRequest("GET", url)
	resp, err := c.makeRequest(req)

	c.unmarshalResponse(ctx, resp, getContributerStatsByRepoResponse)

	return
}

func (c *GithubClientImpl) CheckOrgMembership(ctx context.Context, username string) (isMember bool, err error) {
	url := GITHUB_BASE_URL + "orgs/" + config.Constants.OrgName + "/members/" + username
	req,  err := c.formRequest("GET", url)
	resp, err := c.makeRequest(req)

	// Check if 200 or 404
	// c.unmarshalResponse(ctx, resp, checkOrgMembershipResponse)

	return resp.status_code == 200
}

