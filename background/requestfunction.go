package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	apiKey         = "AIzaSyAa3v8ulaMd6MXQ1oCJDzNCG4pHV6Ms8OU"
	searchEngineID = "227c94475aca5432c"
)

type SearchResult struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}

type FoundContactStruct struct {
	Name        string
	Salutation  string
	Email       []string
	Institution string
	URL         string
}

func makeRequest(payload map[string]string) (*SearchResult, error) {
	// Build the query parameters
	queryParams := url.Values{}
	for key, value := range payload {
		queryParams.Add(key, value)
	}
	queryParams.Add("key", apiKey)
	queryParams.Add("cx", searchEngineID)

	// Send GET request to Google Search API
	resp, err := http.Get("https://www.googleapis.com/customsearch/v1?" + queryParams.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	// Parse JSON response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func buildPayload(query string, start, num int, params map[string]string) map[string]string {
	payload := map[string]string{
		"q":     query,
		"start": fmt.Sprintf("%d", start),
		"num":   fmt.Sprintf("%d", num),
	}
	for key, value := range params {
		payload[key] = value
	}
	return payload
}

func findEmailByNameAndInstitution(urlStr string) ([]string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find and parse email addresses using regex
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	emails := re.FindAllString(string(body), -1)
	return emails, nil
}

// Perform actual requesting and scraping, returning a list of found contacts
func request(firstName, lastName, institution string) []FoundContactStruct {

	searchQuery := firstName + " " + lastName + " " + institution
	totalResults := 20

	var urls []string
	var global_list []FoundContactStruct

	remainder := totalResults % 10
	pages := totalResults / 10
	if remainder > 0 {
		pages++
	}

	for i := 0; i < pages; i++ {
		numResults := 10
		if i == pages-1 && remainder > 0 {
			numResults = remainder
		}
		payload := buildPayload(searchQuery, (i+1)*10, numResults, nil)
		result, err := makeRequest(payload)
		if err != nil {
			log.Fatal(err)
		}
		for _, item := range result.Items {
			urls = append(urls, item.Link)
		}
	}

	// fmt.Println("URLs found:")
	for _, url := range urls {
		if strings.Contains(url, "pdf") {
			continue
		}
		var newresult FoundContactStruct

		emails, err := findEmailByNameAndInstitution(url)
		if err != nil {
			continue
		}

		if len(emails) > 0 && len(emails[0]) <= 100 {
			// fmt.Println(url)
			// fmt.Println(emails)
			newresult.Email = emails
			newresult.URL = url
			newresult.Institution = "N/A"
			newresult.Name = "N/A"
			newresult.Salutation = "N/A"
			global_list = append(global_list, newresult)

		}

	}
	return global_list
}

func main() {
	firstName := "Chris"
	lastName := "Mcdonald"
	institution := "University of Western Australia"

	contacts := request(firstName, lastName, institution)

	fmt.Println("Found Contacts:")
	for _, contact := range contacts {
		fmt.Printf("URL: %s\n", contact.URL)
		fmt.Printf("Email: %v\n", contact.Email)
		fmt.Printf("Institution: %s\n", contact.Institution)
		fmt.Printf("Name: %s\n", contact.Name)
		fmt.Printf("Salutation: %s\n", contact.Salutation)
		fmt.Println("---")
	}
}

// func main() {
// 	firstName := "Chris"
// 	lastName := "Mcdonald"
// 	institution := "University of Western Australia"

// 	fmt.Printf(request(firstName, lastName, institution))

// }

// func main() {

// 	// Use the request function instead of Google Search API to match the desired output
// 	contacts := request("Chris", "Mcdonald", "University of Western Australia")

// 	// Output the found contacts
// 	fmt.Println("Contacts found:")
// 	for _, contact := range contacts {
// 		fmt.Printf("Name: %s\nEmail: %s\nInstitution: %s\nSalutation: %s\nURL: %s\n\n",
// 			contact.Name, contact.Email, contact.Institution, contact.Salutation, contact.URL)
// 	}
// }
