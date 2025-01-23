package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Get(url string, target interface{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error getting response: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
		return nil, err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
		return nil, err
	}

	return target, nil
}
