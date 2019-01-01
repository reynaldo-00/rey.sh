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

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// data := result["data"]
	// viewer := data["viewer"]

	// fmt.Printf("%+v\n", result["data"])
	json.NewEncoder(w).Encode(result["data"])
	log.Println("GET /repos successful")
}

// GetSpotify gets my the song im currently listening to
func GetSpotify(w http.ResponseWriter, req *http.Request) {
	err := godotenv.Load()
	spotifyToken := os.Getenv("SPOTIFY_TOKEN")

	url := "https://api.spotify.com/v1/me/player/currently-playing"

	client := &http.Client{}
	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		log.Fatalln(error)
	}

	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "Bearer "+spotifyToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
	log.Println("GET /spotify successful")
}
