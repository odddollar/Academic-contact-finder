package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/odddollar/CITS3200-Project/global"
)

// Scruct created to store list of URL's found
type ScopusResponse struct {
	SearchResults struct {
		Entry []struct {
			Link []struct {
				Ref  string `json:"@ref"`
				Href string `json:"@href"`
			} `json:"link"`
		} `json:"entry"`
	} `json:"search-results"`
}

func main() {
	// Test details
	firstName := "Chris"
	lastName := "McDonald"
	institution := "University of Western Australia"

	// Build the request URL
	apiUrl := "https://api.elsevier.com/content/search/scopus"
	apiKey := "14c976cb1098fd8fcb283b93d17ed57f"

	// Set up query parameters
	params := url.Values{}
	params.Add("query", fmt.Sprintf("AUTHOR-NAME(%s) AND AFFIL(%s)", firstName+" "+lastName, institution))
	params.Add("apiKey", apiKey)

	// Build the final URL with parameters
	reqUrl := fmt.Sprintf("%s?%s", apiUrl, params.Encode())

	// Create a new request
	resp, err := http.Get(reqUrl)
	if err != nil {
		global.ShowError(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		global.ShowError(err)
	}

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		global.ShowError(err)
	}

	// Parse the JSON response into struct
	var scopusResponse ScopusResponse
	err = json.Unmarshal(body, &scopusResponse)
	if err != nil {
		global.ShowError(err)
	}

	var urls []string

	// Extract URLs from response
	for _, entry := range scopusResponse.SearchResults.Entry {
		for _, link := range entry.Link {
			if link.Ref == "scopus" {
				urls = append(urls, link.Href)
			}
		}
	}

	// Print URLs (for testing)
	fmt.Println("Extracted URLs:")
	for _, url := range urls {
		fmt.Println(url)
	}
}
