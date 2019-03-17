package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reynld/github"
	"github.com/reynld/spotify"
)

// GetServerIsUp '/' engpoint cheks if server is up
func GetServerIsUp(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	message := "GetRequestAt /" + port
	fmt.Println(message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

var port string

func main() {

	r := mux.NewRouter()
	port = ":9001"

	r.HandleFunc("/", GetServerIsUp).Methods("GET")
	r.HandleFunc("/repos", github.GetProjects).Methods("GET")
	r.HandleFunc("/spotify", spotify.GetSpotify).Methods("GET")
	fmt.Println("Server live on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
