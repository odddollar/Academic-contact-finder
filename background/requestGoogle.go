package background

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/odddollar/CITS3200-Project/global"
)

// Struct created to store list of URLs found
type SearchResult struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}

// Perform actual requesting and scraping of google
func requestGoogle(firstName, lastName, institution string) {
	apiKey := "AIzaSyAa3v8ulaMd6MXQ1oCJDzNCG4pHV6Ms8OU"
	searchEngineID := "227c94475aca5432c"
	apiUrl := "https://www.googleapis.com/customsearch/v1?"

	// Setup query data
	searchQuery := fmt.Sprintf("%s %s %s", firstName, lastName, institution)

	// Build the query parameters
	params := url.Values{}
	params.Add("q", searchQuery)
	params.Add("key", apiKey)
	params.Add("cx", searchEngineID)

	// Build final url with parameters
	reqUrl := fmt.Sprintf("%s%s", apiUrl, params.Encode())

	// Send GET request to Google Search API
	resp, err := http.Get(reqUrl)
	if err != nil {
		global.ShowError(err)
	}
	defer resp.Body.Close()

	// Check if response is successful
	if resp.StatusCode != http.StatusOK {
		global.ShowError(errors.New("Bad http response"))
	}

	// Parse JSON response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.ShowError(err)
	}

	// Parse JSON response into struct
	var result SearchResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		global.ShowError(err)
	}
	fmt.Println(result)

	var results []global.FoundContactStruct

	global.AllFoundContacts = append(global.AllFoundContacts, results...)
}
