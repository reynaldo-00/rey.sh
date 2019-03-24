package spotify

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// getAccessToken Makes call to spotify api to get a refreshed token
func getAccessToken() string {
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
	spotifyToken := getAccessToken()

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

	json.NewEncoder(w).Encode(result)
	log.Println("GET /spotify successful")
}
