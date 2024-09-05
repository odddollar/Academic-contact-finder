package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Scruct created to store list of URL's found
type ScopusResponse struct {
	SearchResults struct {
		Entry []struct {
			Link []struct {
				Rel  string `json:"@ref"`
				Href string `json:"@href"`
			} `json:"link"`
		} `json:"entry"`
	} `json:"search-results"`
}

func main() {

	//test details
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
	//params.Add("view", "complete")

	// Build the final URL with parameters
	reqUrl := fmt.Sprintf("%s?%s", apiUrl, params.Encode())

	// Create a new HTTP GET request
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	fmt.Println(string(body)) //FOR TESTING

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Received status code %d", resp.StatusCode)
	}

	// Parse the JSON response into a Go struct
	var scopusResponse ScopusResponse
	err = json.Unmarshal(body, &scopusResponse)
	if err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	// Create a Slice to hold the URLs
	var urls []string

	// extract the URLs from the response
	for _, entry := range scopusResponse.SearchResults.Entry {
		for _, link := range entry.Link {
			if link.Rel == "scopus" {
				urls = append(urls, link.Href)
			}
		}
	}

	// Print the URLs (for testing)
	fmt.Println("Extracted URLs:")
	for _, url := range urls {
		fmt.Println(url)
	}

}
