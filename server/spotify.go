package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// TokenResponse struct to mirror refresh token response
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// GetAccessToken Makes call to spotify api to get a refreshed token
func GetAccessToken() string {
	godotenv.Load()
	refereshURL := "https://accounts.spotify.com/api/token"
	spotifyRefreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")
	spotifyClientID := os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", spotifyRefreshToken)
	data.Set("client_id", spotifyClientID)
	data.Set("client_secret", spotifyClientSecret)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", refereshURL, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)

	var result TokenResponse
	json.NewDecoder(resp.Body).Decode(&result)

	// fmt.Println(result["access_token"])
	return result.AccessToken
}

// GetSpotify gets my the song im currently listening to
func GetSpotify(w http.ResponseWriter, req *http.Request) {
	spotifyToken := GetAccessToken()
	// fmt.Println(spotifyToken)

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
