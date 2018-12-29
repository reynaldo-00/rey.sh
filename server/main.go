package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var port string

func main() {

	r := mux.NewRouter()
	port = ":9001"

	r.HandleFunc("/", GetServerIsUp).Methods("GET")
	fmt.Println("Server live on port " + port)
	http.ListenAndServe(port, r)
}
