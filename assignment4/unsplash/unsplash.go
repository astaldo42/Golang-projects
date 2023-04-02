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


