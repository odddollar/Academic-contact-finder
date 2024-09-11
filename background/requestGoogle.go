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

	// Build the query parameters
	params := url.Values{}
	params.Add("q", searchQuery)
	params.Add("key", apiKey)
	params.Add("cx", searchEngineID)

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
	var urls []string
	for _, i := range googleResults.Items {
		urls = append(urls, i.Link)
	}

	var results []global.FoundContactStruct

	// Iterate through urls
	for _, i := range urls {
		// Scrape data from current url
		r := scrapeSite(i, ctx)

		fmt.Println(r)
	}

	global.AllFoundContacts = append(global.AllFoundContacts, results...)
}

// Take url and chromedp context and scrape data from site
func scrapeSite(u string, ctx context.Context) []global.FoundContactStruct {
	// Create variables to store the page's HTML and data found
	var htmlContent string
	var toReturn []global.FoundContactStruct

	// Visit the webpage and get the HTML content
	err := chromedp.Run(ctx,
		chromedp.Navigate(u),
		chromedp.Sleep(1*time.Second), // Adding sleep for reliability
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		global.ShowError(err)
	}

	// Parse email regex
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find first email
	foundEmail := re.FindString(htmlContent)

	fmt.Println(foundEmail)

	return toReturn
}
