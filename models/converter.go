package models

import (
	"github.com/davecgh/go-spew/spew"
)

func ConvertOrgStatsToOrgStatsByUser(orgStats OrgStats) (orgStatsByUser *OrgStatsByUser, err error) {
	result := make(OrgStatsByUser)
	for repo_name, repo_stats := range orgStats {
		for _, repo_stat := range repo_stats {
			userStatsByRepo := make(UserStatsByRepo)
			github_username := repo_stat.Author.Login
			result[github_username] = &userStatsByRepo

			total_commits   := 0
			total_additions := 0
			total_deletions := 0
			for _, week := range repo_stat.Weeks {
				total_commits   += week.Commits
				total_additions += week.Additions
				total_deletions += week.Deletions
			}

			aggregateRepoStats := AggregateRepoStats{
				TotalCommits:   total_commits,
				TotalAdditions: total_additions,
				TotalDeletions: total_deletions,
				Weeks: 					repo_stat.Weeks,
			}

			userStatsByRepo[repo_name] = &aggregateRepoStats
			spew.Dump(userStatsByRepo)
		}
	}

	return &result, nil
}