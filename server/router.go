package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	pinnedRepo := GQLQuery{
		Query: "query {\n  viewer {\n    pinnedRepositories(first: 10){\n      totalCount\n      nodes {\n        id\n        url\n        pushedAt\n        shortDescriptionHTML\n        defaultBranchRef {\n          target {\n            ...on Commit {\n              history(first:10) {\n                totalCount\n              }\n            }\n          }\n        }\n        repositoryTopics(first:10) {\n          edges {\n            node {\n              id\n              topic {\n                name\n              }\n            }\n          }\n        }\n        languages(first: 10) {\n          nodes {\n            color\n            id\n            name\n          }\n        }\n        primaryLanguage {\n          color\n          id\n          name\n        }\n      }\n    }\n  }\n}",
	}

	pinquery, _ := json.Marshal(pinnedRepo)

	client := &http.Client{}

	req, error := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(pinquery))
	if error != nil {
		log.Fatalln(error)
	}
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", "bearer githubkey")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	// log.Println(result)
	// log.Println(result["data"])
	json.NewEncoder(w).Encode(result["data"])
	log.Println("GET /repos successful")
}
