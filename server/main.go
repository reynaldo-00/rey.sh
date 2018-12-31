package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string

func main() {

	r := mux.NewRouter()
	port = ":9001"

	r.HandleFunc("/", GetServerIsUp).Methods("GET")
	r.HandleFunc("/repos", GetProjects).Methods("GET")
	r.HandleFunc("/spotify", GetSpotify).Methods("GET")
	fmt.Println("Server live on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
