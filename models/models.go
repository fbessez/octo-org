package models

type OrgStats map[string][]*Contributor

// github_name => UserStatsByRepo
type OrgStatsByUser map[string]*UserStatsByRepo

// repo_name => AggregateRepoStats
type UserStatsByRepo map[string]*AggregateRepoStats

type AggregateRepoStats struct {
	TotalCommits    int
	TotalAdditions  int
	TotalDeletions  int
	Weeks           []*WeekContribution
}

type UserCommits struct {
	GithubUsername string
	TotalCommits   int
}