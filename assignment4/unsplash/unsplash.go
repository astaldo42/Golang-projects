package unsplash

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RandomPhoto() (string, error) { // thats gonna give me the url of photo
	url := "https://api.unsplash.com/photos/random?client_id=" + access
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	urls, ok := data["urls"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("no urls found")
	}

	photo, ok := urls["small"].(string)
	if !ok {
		return "", fmt.Errorf("url is not found")
	}

	return photo, nil
}

func GetRandomPhotoURL() (string, error) {
	// Make a GET request to the Unsplash API
	resp, err := http.Get("https://api.unsplash.com/photos/random?client_id=ci2S282CMJqI5xd__z7Si26fQixbOjSLHOBv4-YMg6o")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode the JSON response
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	// Check for any errors in the API response
	if _, ok := data["errors"]; ok {
		return "", fmt.Errorf("API returned an error: %v", data["errors"])
	}

	// Extract the URL of the photo from the API response
	photoURL := data["urls"].(map[string]interface{})["regular"].(string)

	return photoURL, nil
}
