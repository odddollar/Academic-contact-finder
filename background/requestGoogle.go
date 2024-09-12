package background

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
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

	// Create new Chrome browser context with options to disable headless mode
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),        // Disable headless mode
		chromedp.Flag("disable-gpu", false),     // Enable GPU usage
		chromedp.Flag("start-maximized", false), // Start maximized
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create chromedp context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Setup query data
	searchQuery := fmt.Sprintf("%s %s %s", firstName, lastName, institution)

	// Calculate number of results and pages to ask from google
	totalResults := 20
	remainder := totalResults % 10
	pages := totalResults / 10
	if remainder > 0 {
		pages++
	}

	var urls []string

	// Iterate through number of pages on google
	for i := 0; i < pages; i++ {
		numResults := 10
		if i == pages-1 && remainder > 0 {
			numResults = remainder
		}

		// Build the query parameters
		params := url.Values{}
		params.Add("key", apiKey)
		params.Add("cx", searchEngineID)
		params.Add("q", searchQuery)
		params.Add("start", fmt.Sprintf("%d", i*10))
		params.Add("num", fmt.Sprintf("%d", numResults))

		// Build final url with parameters
		reqUrl := fmt.Sprintf("%s%s", googleApiUrl, params.Encode())

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

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			global.ShowError(err)
		}

		// Parse JSON response into struct
		var googleResults SearchResult
		err = json.Unmarshal(body, &googleResults)
		if err != nil {
			global.ShowError(err)
		}

		// Get urls
		for _, i := range googleResults.Items {
			urls = append(urls, i.Link)
		}
	}

	// Iterate through urls
	for _, i := range urls {
		// Skip pdfs
		if strings.Contains(i, ".pdf") {
			continue
		}

		// Scrape data from current url
		r, valid := scrapeSite(i, ctx, firstName, lastName, institution)

		// Append found valid result directly to array
		if valid {
			fmt.Println(r)
			global.AllFoundContacts = append(global.AllFoundContacts, r)
		}
	}
}

// Take url and chromedp context and scrape data from site.
// Each site will return maximum one found result, with bool indicating if result
// is valid (i.e. has an email been found)
func scrapeSite(u string, ctx context.Context, firstName, lastName, institution string) (global.FoundContactStruct, bool) {
	// Store page's html content
	var htmlContent string

	// Visit the webpage and get the HTML content
	err := chromedp.Run(ctx,
		chromedp.Navigate(u),
		chromedp.Sleep(1*time.Second), // Adding sleep for reliability
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		global.ShowError(err)
	}

	// Convert html to lowercase
	htmlContentLower := strings.ToLower(htmlContent)

	// Parse email regex
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find all emails and filter down to matching ones
	foundEmails := re.FindAllString(htmlContentLower, -1)
	var matchingEmails []string
	for _, i := range foundEmails {
		// Extract the part before the "@" symbol and convert to lowercase
		localPart := strings.ToLower(strings.Split(i, "@")[0])

		// Check if local part contains first or last name
		if strings.Contains(localPart, firstName) && firstName != "" {
			matchingEmails = append(matchingEmails, i)
		} else if strings.Contains(localPart, lastName) && lastName != "" {
			matchingEmails = append(matchingEmails, i)
		}
	}

	// If more than one email found then just use first
	var email string
	if len(matchingEmails) >= 1 {
		email = matchingEmails[0]
	}

	fmt.Println(matchingEmails, email)

	// Set first name, last name, or institution to N/A if one isn't found
	if !strings.Contains(htmlContentLower, firstName) || firstName == "" {
		firstName = "N/A"
	} else {
		firstName = findExactMatch(htmlContent, htmlContentLower, firstName)
	}
	if !strings.Contains(htmlContentLower, lastName) {
		lastName = "N/A"
	} else {
		lastName = findExactMatch(htmlContent, htmlContentLower, lastName)
	}
	if !strings.Contains(htmlContentLower, institution) || institution == "" {
		institution = "N/A"
	} else {
		institution = findExactMatch(htmlContent, htmlContentLower, institution)
	}

	// Format results to correct structure
	up, _ := url.Parse(u)
	result := global.FoundContactStruct{
		FirstName:   firstName,
		LastName:    lastName,
		Salutation:  "",
		Email:       email,
		Institution: institution,
		URL:         up,
	}

	// "Invalid" result if no email found
	if email == "" {
		return result, false
	}
	return result, true
}

// Searches for case-insensitive occurrence of short string within long string
// and returns exact case-sensitive substring from original string
func findExactMatch(original, originalLower, toFind string) string {
	// Find starting index of desired string in original lower string (case insensitive)
	index := strings.Index(originalLower, toFind)

	// Failsafe to prevent program crashing
	if index == -1 {
		return "N/A"
	}

	// Extract exact substring from original string using index and length of desired string
	return original[index : index+len(toFind)]
}
