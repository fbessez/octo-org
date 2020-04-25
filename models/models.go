package models

type OrgStats map[string][]*Contributor

// github_name => UserStatsByRepo
type OrgStatsByUser map[string]*UserStatsByRepo

// repo_name => AggregateRepoStats
type UserStatsByRepo map[string]*AggregateRepoStats

type AggregateRepoStats struct {
	TotalCommits    int                 `json:"total_commits`
	TotalAdditions  int                 `json:"total_additions"`
	TotalDeletions  int                 `json:"total_deletions"`
	Weeks           []*WeekContribution `json:"weeks"`
}

type UserCommits struct {
	GithubUsername string  `json:"github_username"`
	TotalCommits   int     `json:"total_commits"`
	TotalAdditions int     `json:"total_additions"`
	TotalDeletions int     `json:"total_deletions"`
}