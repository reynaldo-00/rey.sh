package main

import (
	"fmt"
	"net/http"
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
