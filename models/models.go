package models

type Repository struct {
	Id 		 					int 	 `json:"id"`
	NodeId 					string `json:"node_id"`
	Name   					string `json:"name"`
	FullName 				string `json:"full_name"`
	Owner 		   		*Owner `json:"owner"`
	ContributorsUrl string `json:"contributors_url`
	Archived        bool   `json:"archived"`
}

type Owner struct {
	Id     int    `json:"id"`
	NodeId string `json:"node_id"`
	Login  string `json:"login"`
	Url    string `json:"url"`
}

type Contributor struct {
	TotalCommits int                 `json:"total"`
	Weeks        []*WeekContribution `json:"weeks"`
	Author       *Owner              `json:"author"`
}

type WeekContribution struct {
	Week      int `json:"w"`
	Additions int `json:"a"`
	Deletions int `json:'d"`
	Commits   int `json:"c"`
}