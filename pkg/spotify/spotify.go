// pkg/spotify/spotify.go
package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Item struct {
	Album struct {
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	} `json:"album"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Popularity  int    `json:"popularity"`
	Name        string `json:"name"`
	ExternalIDs struct {
		ISRC string `json:"isrc"`
	} `json:"external_ids"`
}

type SpotifyResponse struct {
	Tracks struct {
		Items []Item `json:"items"`
	} `json:"tracks"`
}

// GetTrackByISRC fetches metadata from Spotify for a given ISRC
func GetTrackByISRC(isrc string) (*SpotifyResponse, error) {
	// Replace "YOUR_SPOTIFY_API_KEY" with your actual Spotify API key
	spotifyAPIURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=isrc:%s&amp;type=track", isrc)
	req, err := http.NewRequest("GET", spotifyAPIURL, nil)
	if err != nil {
		return nil, err
	}

	token, err := getSpotifyAccessToken()

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for errors in the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Spotify API returned non-OK status: %s", resp.Status)
	}

	// Parse the response body into SpotifyResponse struct
	var spotifyResponse SpotifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&spotifyResponse); err != nil {
		return nil, err
	}

	return &spotifyResponse, nil
}

func getSpotifyAccessToken() (string, error) {
	tokenURL := "https://accounts.spotify.com/api/token"
	clientID := "09ebbaeee98247e2b40fb4293da990f4"
	clientSecret := "1689b0d4dccf442690ae1022165b9c2a"
	grantType := "client_credentials"

	// Create a new HTTP request to get the access token
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(
		"grant_type="+grantType+
			"&client_id="+clientID+
			"&client_secret="+clientSecret,
	))
	if err != nil {
		return "", err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for errors in the response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Spotify API returned non-OK status during token retrieval: %s", resp.Status)
	}

	// Parse the response body to get the access token
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
