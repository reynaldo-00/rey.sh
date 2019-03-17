package github

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// GQLQuery struct to hold query string
type GQLQuery struct {
	// Query in string form
	Query string `json:"query"`
}

// GithubResp response struct to store response from github api
type GithubResp struct {
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

// GetProjects gets pinned repositories of my github profile
func GetProjects(w http.ResponseWriter, req *http.Request) {
	err := godotenv.Load()
	githubKey := os.Getenv("GIT_API_KEY")
	pinnedRepo := GQLQuery{
		Query: `query {
			viewer {
			  pinnedRepositories(first: 10){
				totalCount
				nodes {
					id
					name
				  url
				  pushedAt
				  shortDescriptionHTML
				  defaultBranchRef {
					target {
					  ...on Commit {
						history(first:10) {
						  totalCount
						}
					  }
					}
				  }
				  repositoryTopics(first:10) {
					edges {
					  node {
						id
						topic {
						  name
						}
					  }
					}
				  }
				  languages(first: 10) {
					nodes {
					  color
					  id
					  name
					}
				  }
				  primaryLanguage {
					color
					id
					name
				  }
				}
			  }
			}
		  }
		`}
	pinquery, _ := json.Marshal(pinnedRepo)

	client := &http.Client{}
	req, error := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(pinquery))
	if error != nil {
		log.Fatalln(error)
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "bearer "+githubKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var result GithubResp
	json.NewDecoder(resp.Body).Decode(&result)
	// log.Printf("Recieved: %v\n", result)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(result.Data.Viewer.PinnedRepositories.Nodes)
	log.Println("GET /repos successful")
}
