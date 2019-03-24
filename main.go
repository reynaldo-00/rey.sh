package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/reynld/rey.sh/pkg/github"
	"github.com/reynld/rey.sh/pkg/spotify"
)

// GetServerIsUp '/' engpoint cheks if server is up
func GetServerIsUp(w http.ResponseWriter, req *http.Request) {
	message := "GetRequestAt /" + port
	fmt.Println(message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

var port string

func main() {
	godotenv.Load()
	r := mux.NewRouter()
	port = ":9001"

	r.HandleFunc("/", GetServerIsUp).Methods("GET")
	r.HandleFunc("/repos", github.GetProjects).Methods("GET")
	r.HandleFunc("/spotify", spotify.GetSpotify).Methods("GET")
	fmt.Println("Server live on port " + port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS()(r)))
}
