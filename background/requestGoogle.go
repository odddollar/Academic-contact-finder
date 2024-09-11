package background

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/odddollar/CITS3200-Project/global"
	"github.com/odddollar/CITS3200-Project/widgets"
)

// Initiate api requesting and scraping, then update results
func _() {
	// Ensure API key is present and valid
	// Run here again, as if cancel clicked initially then still no api key
	// This will appear until a valid key is entered every time run is clicked
	// and will not progress running any futher
	// if !PresentAPIKey() || !ValidAPIKey() {
	// 	UpdateAPIKey()
	// 	return
	// }

	// Get data from entry boxes
	firstName := global.Ui.FirstName.Text
	lastName := global.Ui.LastName.Text
	institution := global.Ui.Institution.Text

	// Create and show loading bar
	loading := infiniteLoad()
	loading.Show()

	// Make request and get results
	global.AllFoundContacts = request(firstName, lastName, institution)

	// Hide loading bar
	loading.Hide()

	// Enable email all button
	global.Ui.EmailAll.Enable()

	// Update number of results found
	global.Ui.NumResults.Text = fmt.Sprintf("Found %d results", len(global.AllFoundContacts))
	global.Ui.NumResults.Refresh()

	// Iterate through returned results and update UI
	global.Ui.Output.RemoveAll()
	for i := 0; i < len(global.AllFoundContacts); i++ {
		global.Ui.Output.Add(widgets.NewFoundContact(global.AllFoundContacts[i]))
	}
}

const (
	apiKey         = "AIzaSyAa3v8ulaMd6MXQ1oCJDzNCG4pHV6Ms8OU"
	searchEngineID = "227c94475aca5432c"
)

type SearchResult struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
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

func findEmailByNameAndInstitution(urlStr string, name string) ([]string, error) {
	var emails []string

	// Perform HTTP GET request
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find and parse email addresses using regex
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	foundEmails := re.FindAllString(string(body), -1)

	// Convert name to lowercase for case-insensitive comparison
	name = strings.ToLower(name)

	// Loop through found emails and match against the provided name
	for _, email := range foundEmails {
		// Extract the part before the '@' symbol and convert to lowercase
		localPart := strings.ToLower(strings.Split(email, "@")[0])

		// Check if the local part contains the name
		if strings.Contains(localPart, name) {
			emails = append(emails, email)
		}
	}

	return emails, nil
}

func findInstitutionandName(urlStr string, institution string, name string) (string, string, error) {
	var institutionresult string
	var name_result string
	institutionresult = institution
	name_result = name

	// Perform HTTP GET request
	resp, err := http.Get(urlStr)
	if err != nil {
		return "N/A", "N/A", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "N/A", "N/A", err
	}

	// Convert body to string and make it lowercase for case-insensitive search
	bodyStr := strings.ToLower(string(body))
	// Convert institution to lowercase for case-insensitive comparison
	institution = strings.ToLower(institution)

	name = strings.ToLower(institution)

	// Check if the institution string is in the HTML body
	if !strings.Contains(bodyStr, institution) {
		institutionresult = "N/A"
	}

	if !strings.Contains(bodyStr, name) {
		name_result = "N/A"
	}

	return institutionresult, name_result, nil
}

// Perform actual requesting and scraping, returning a list of found contacts
func request(firstName, lastName, institution string) []global.FoundContactStruct {

	searchQuery := firstName + " " + lastName + " " + institution
	totalResults := 20
	name := firstName + " " + lastName

	var urls []string

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
	for _, urlString := range urls {
		if strings.Contains(urlString, "pdf") {
			continue
		}
		var newresult global.FoundContactStruct

		emails, err := findEmailByNameAndInstitution(urlString, firstName)
		if err != nil {
			continue
		}
		institutionresult, name_result, err := findInstitutionandName(urlString, institution, name)
		if err != nil {
			continue
		}
		if len(emails) > 0 && len(emails[0]) <= 100 {
			// fmt.Println(url)
			// fmt.Println(emails)
			newresult.Email = emails[0]
			parsedURL, _ := url.Parse(urlString)
			newresult.URL = parsedURL
			newresult.Institution = institutionresult
			newresult.FirstName = name_result
			newresult.Salutation = "N/A"
			global.AllFoundContacts = append(global.AllFoundContacts, newresult)

		}

	}
	return global.AllFoundContacts
}
