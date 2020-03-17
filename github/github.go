package github

import (
	"context"
	"github.com/fbessez/orgContributions/models"
)

type GithubClient interface {
	GetOrg(ctx context.Context) (response *models.GetOrgResponse, err error)
	GetUser(ctx context.Context) (response *models.GetUserResponse, err error)
	GetReposByOrg(ctx context.Context) (response *models.GetReposByOrgResponse, err error)
	CheckOrgMembership(ctx context.Context, org) (response *models.CheckOrgMembershipResponse, err error)
	GetContributerStatsByRepo(ctx context.Context) (response *models.GetContributerStatsByRepoResponse, err error)
}