package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DatamuseResult struct {
	Word  string `json:"word"`
	Score int    `json:"score"`
}

func FetchWordsFromDatamuse(word string) ([]DatamuseResult, error) {
	// Construct the URL for the API request
	url := fmt.Sprintf("%s%s", DATAMUSE_API, word)

	// Send the GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from Datamuse API: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned an error: %s", resp.Status)
	}

	// Parse the JSON response into a slice of Word structs
	var words []DatamuseResult
	if err := json.Unmarshal(body, &words); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	// Return the list of words
	return words, nil
}
