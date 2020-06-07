package models

type CheckOrgMembershipResponse struct {
	// 200 if success
	// 404 if not
}

type GetOrgResponse struct {
	Id          int    `json:"id"`
	NodeId      string `json:"node_id"`
	Login       string `json:"login"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	ReposUrl    string `json:"repos_url"`
	MembersUrl  string `json:"members_url"`
	AvatarUrl   string `json:"avatar_url"`
}

type GetUserResponse struct {
	Id               int    `json:"id"`
	NodeId           string `json:"node_id"`
	Login            string `json:"login"`
	Name             string `json:"name"`
	Bio              string `json:"bio"`
	Email            string `json:"email"`
	Location         string `json:"location"`
	Company          string `json:"company"`
	Blog             string `json:"blog"`
	Type             string `json:"type"`
	Url              string `json:"url"`
	AvatarUrl        string `json:"avatar_url"`
	HtmlUrl          string `json:"html_url"`
	OrganizationsUrl string `json:"organizations_url"`
	ReposUrl         string `json:"repos_url"`
	Hireable         bool   `json:"hireable"`
	SiteAdmim        bool   `json:"site_admin"`
}

type GetReposByOrgResponse struct {
	Repos []*Repository
}

type GetContributerStatsByRepoResponse struct {
	Contributors []*Contributor
}
