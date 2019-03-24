package github

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// GetProjects gets pinned repositories of my github profile
func GetProjects(w http.ResponseWriter, req *http.Request) {
	githubKey := os.Getenv("GIT_API_KEY")
	pinnedRepo := GQLQuery{Query: githubQuery}
	pinquery, _ := json.Marshal(pinnedRepo)

	client := &http.Client{}
	req, error := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(pinquery))
	if error != nil {
		log.Fatalln(error)
	}

	req.Header.Add("Authorization", "bearer "+githubKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var result GithubResp
	json.NewDecoder(resp.Body).Decode(&result)

	json.NewEncoder(w).Encode(result.Data.Viewer.PinnedRepositories.Nodes)
	log.Println("GET /repos successful")
}
