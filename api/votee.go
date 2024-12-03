package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var vowelList = []string{"a", "e", "i", "o", "u", "y"}

// otherLetterList ordered by most common letters decreasing
var otherLetterList = []string{
	"t", "n", "s", "h", "r", "d", "l", "c", "m", "f", "p", "b", "v", "k", "j", "x", "q", "z", "w", "g",
}

type VoteeResult struct {
	Slot   int    `json:"slot"`
	Guess  string `json:"guess"`
	Result string `json:"result"`
}

func GuessWord(mode string, size int, seed int) (string, error) {
	// check if word contain vowels
	word, correctCount, err := findCorrectPosition(vowelList, size, mode, seed)

	// fetch the list of possible match from datamuse
	var datamuseResult []DatamuseResult
	if correctCount > 0 {
		datamuseResult, err = FetchWordsFromDatamuse(word)
		if err != nil {
			return "", err
		}
	} else { // case is word does not contain any vowels, check with other letters, it has early stop
		word, correctCount, err = findCorrectPosition(vowelList, size, mode, seed)
	}

	// after that, guess with this list
	res, err := guessWithResultFromDatamuse(datamuseResult, size, mode, seed)
	if err != nil {
		return "", err
	}

	return res, nil
}

func guessWithResultFromDatamuse(datamuseResult []DatamuseResult, size int, mode string, seed int) (string, error) {
	for _, result := range datamuseResult {
		url := fmt.Sprintf("%sdaily?guess=%s&size=%d", VOTEE_API, result.Word, size)

		if mode == RANDOM {
			url = fmt.Sprintf("%s&seed=%d", url, seed)
		}

		voteeResult, err := guessWord(url)
		if err != nil {
			return "", err
		}

		if checkResult(voteeResult) {
			return result.Word, nil
		}
	}
	return "", nil
}

func findCorrectPosition(letterList []string, size int, mode string, seed int) (string, int, error) {
	word := populate("?", size)

	correctCount := 0
	// check all vowels and get its correct position
	for _, letter := range letterList {

		//early stop in case we have to check with otherLetterList
		if correctCount >= size/2 && len(letterList) > 6 {
			return word, correctCount, nil
		}

		url := fmt.Sprintf("%sdaily?guess=%s&size=%d", VOTEE_API, populate(letter, size), size)
		if mode == RANDOM {
			url = fmt.Sprintf("%s&seed=%d", url, seed)
		}

		voteeResult, err := guessWord(url)
		if err != nil {
			return "", 0, err
		}

		for _, r := range voteeResult {
			if r.Result == CORRECT {
				word = word[:r.Slot] + letter + word[r.Slot+1:]
				correctCount++
			}
		}
	}

	return word, correctCount, nil
}

func populate(letter string, size int) string {
	res := ""
	for i := 0; i < size; i++ {
		res += letter
	}
	return res
}

func checkResult(res []VoteeResult) bool {
	for _, r := range res {
		if r.Result != CORRECT {
			return false
		}
	}
	return true
}

func guessWord(url string) ([]VoteeResult, error) {
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
	var result []VoteeResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	// Return the list of result
	return result, nil
}
