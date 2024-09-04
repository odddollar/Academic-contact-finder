package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type ScopusResponse struct {
	SearchResults struct {
		Entry []struct {
			Title string `json:"dc:title"`
			Link  []struct {
				Rel  string `json:"@ref"`
				Href string `json:"@href"`
			} `json:"link"`
			PublicationDate string `json:"prism:coverDate"`
		} `json:"entry"`
	} `json:"search-results"`
}

func main() {

	//test details
	firstName := "Chris"
	lastName := "McDonald"
	institution := "UWA"

	// Build the request URL
	apiUrl := "https://api.elsevier.com/content/search/scopus"
	apiKey := "14c976cb1098fd8fcb283b93d17ed57f"

	// Set up query parameters
	params := url.Values{}
	params.Add("query", fmt.Sprintf("AUTHFIRST(%s) AND AUTHLAST(%s) AND AFFIL(%s)", firstName, lastName, institution))
	params.Add("apiKey", apiKey)

	// Build the final URL with parameters
	reqUrl := fmt.Sprintf("%s?%s", apiUrl, params.Encode())

	// Create a new HTTP GET request
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

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

	// Print the titles, URLs, and publication dates
	for _, entry := range scopusResponse.SearchResults.Entry {
		var paperUrl string
		for _, link := range entry.Link {
			if link.Rel == "scopus" { // Assuming "scopus" relation contains the paper URL
				paperUrl = link.Href
				break
			}
		}
		fmt.Printf("Title: %s\nURL: %s\nPublication Date: %s\n\n", entry.Title, paperUrl, entry.PublicationDate)
	}
}
