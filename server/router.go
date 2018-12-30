package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// GQLQuery struct to hold query string
type GQLQuery struct {
	// Query in string form
	Query string `json:"query"`
}

// GetServerIsUp '/' engpoint cheks if server is up
func GetServerIsUp(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	message := "GetRequestAt /" + port
	fmt.Println(message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// GetProjects gets pinned repositories of my profile
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
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	// log.Println(result)
	// log.Println(result["data"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result["data"])
	log.Println("GET /repos successful")
}
