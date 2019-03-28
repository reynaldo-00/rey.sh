package main

import (
	"net/http"
)

//Handler function for now.sh
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server is live"))
}
