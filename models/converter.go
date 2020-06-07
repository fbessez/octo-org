package models

func ConvertOrgStatsToOrgStatsByUser(orgStats OrgStats) (orgStatsByUser *OrgStatsByUser, err error) {
	result := make(OrgStatsByUser)
	for repo_name, repo_stats := range orgStats {
		for _, repo_stat := range repo_stats {
			github_username := repo_stat.Author.Login
			userStatsByRepo := make(UserStatsByRepo)

			total_commits := 0
			total_additions := 0
			total_deletions := 0
			for _, week := range repo_stat.Weeks {
				total_commits += week.Commits
				total_additions += week.Additions
				total_deletions += week.Deletions
			}

			aggregateRepoStats := AggregateRepoStats{
				TotalCommits:   total_commits,
				TotalAdditions: total_additions,
				TotalDeletions: total_deletions,
				Weeks:          repo_stat.Weeks,
			}

			userStatsByRepo[repo_name] = &aggregateRepoStats

			if result[github_username] == nil {
				result[github_username] = &userStatsByRepo
			}
			for repo_name, repo_stats := range *result[github_username] {
				userStatsByRepo[repo_name] = repo_stats
			}

			result[github_username] = &userStatsByRepo
		}
	}

	return &result, nil
}
