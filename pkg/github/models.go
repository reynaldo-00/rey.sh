package github

import "time"

// GQLQuery struct to hold query string
type GQLQuery struct {
	// Query in string form
	Query string `json:"query"`
}

// GithubResp response struct to store response from github api
type Resp struct {
	Data struct {
		Viewer struct {
			PinnedRepositories struct {
				TotalCount int `json:"totalCount"`
				Nodes      []struct {
					ID                   string    `json:"id"`
					Name                 string    `json:"name"`
					URL                  string    `json:"url"`
					PushedAt             time.Time `json:"pushedAt"`
					ShortDescriptionHTML string    `json:"shortDescriptionHTML"`
					DefaultBranchRef     struct {
						Target struct {
							History struct {
								TotalCount int `json:"totalCount"`
							} `json:"history"`
						} `json:"target"`
					} `json:"defaultBranchRef"`
					RepositoryTopics struct {
						Edges []struct {
							Node struct {
								ID    string `json:"id"`
								Topic struct {
									Name string `json:"name"`
								} `json:"topic"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"repositoryTopics"`
					Languages struct {
						Nodes []struct {
							Color string `json:"color"`
							ID    string `json:"id"`
							Name  string `json:"name"`
						} `json:"nodes"`
					} `json:"languages"`
					PrimaryLanguage struct {
						Color string `json:"color"`
						ID    string `json:"id"`
						Name  string `json:"name"`
					} `json:"primaryLanguage"`
				} `json:"nodes"`
			} `json:"pinnedRepositories"`
		} `json:"viewer"`
	} `json:"data"`
}
