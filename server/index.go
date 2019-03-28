package main

import (
	"net/http"
)

//Handler function for now.sh
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("server is live"))
}
