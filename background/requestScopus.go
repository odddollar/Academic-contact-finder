package background

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/odddollar/CITS3200-Project/global"
)

// Struct created to store list of URLs found
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

// Perform actual requesting and scraping, returning list of found contacts
func requestScopus(firstName, lastName, institution string) {
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

	// Build the request URL
	apiUrl := "https://api.elsevier.com/content/search/scopus"
	apiKey := global.A.Preferences().String("Scopus_API_key")

	// Set up query based on whether affiliation provided
	var query string
	if institution == "" {
		query = fmt.Sprintf("AUTHOR-NAME(%s)", firstName+" "+lastName)
	} else {
		query = fmt.Sprintf("AUTHOR-NAME(%s) AND AFFIL(%s)", firstName+" "+lastName, institution)
	}

	// Set up query parameters
	params := url.Values{}
	params.Add("query", query)
	params.Add("apiKey", apiKey)

	// Build final URL with parameters
	reqUrl := fmt.Sprintf("%s?%s", apiUrl, params.Encode())

	// Create a new request
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

	results := []global.FoundContactStruct{}

	// Iterate through urls
	for _, i := range urls {
		// Scrape data from current url
		r := scrapeScopus(i, ctx)

		// Iterate through results from url
		for _, j := range r {
			// Check if result qualifies as match
			if strings.ToLower(j.FirstName) == firstName || strings.ToLower(j.LastName) == lastName {
				results = append(results, j)
			}
		}
	}

	global.AllFoundContacts = append(global.AllFoundContacts, results...)
}

// Take url and chromedp context and scrape data from scopus
func scrapeScopus(u string, ctx context.Context) []global.FoundContactStruct {
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

	// Parse the HTML with goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		global.ShowError(err)
	}

	// Create affiliation map
	affiliations := generateAffiliationMap(doc)

	// Find all <li> elements
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		// Find all <a> elements with href attributes
		s.Find("a[href]").Each(func(j int, a *goquery.Selection) {
			href, exists := a.Attr("href")
			if exists && strings.HasPrefix(href, "mailto:") {
				// Find the <span> within the <button> to get the name
				name := s.Find("button span").Text()

				// Prevent error if person only has one name
				var firstName, lastName string
				if len(strings.Split(name, ", ")) > 1 {
					firstName = strings.Split(name, ", ")[1]
					lastName = strings.Split(name, ", ")[0]
				} else {
					firstName = name
					lastName = ""
				}

				// Find the <sup> to get affiliation link
				affiliationLink := s.Find("sup").Text()
				affiliation := affiliations[affiliationLink]
				if affiliation == "" {
					affiliation = affiliations["a"]
				}

				// Format results to correct structure
				up, _ := url.Parse(u)
				toReturn = append(toReturn, global.FoundContactStruct{
					FirstName:   firstName,
					LastName:    lastName,
					Salutation:  "",          // Salutation not provided by scopus
					Email:       href[7:],    // Remove "mailto:"
					Institution: affiliation, // Get affiliation from map
					URL:         up,          // Parsed url as source
				})
			}
		})
	})

	return toReturn
}

// Take document and generate map of superscript identifiers to institutions.
// E.g. "ᵃ University of Western Australia", will be added to the map as
// map[a:University of Western Australia]
func generateAffiliationMap(doc *goquery.Document) map[string]string {
	// Make empty map
	toReturn := make(map[string]string)

	// Find <div> that contains all affiliations. Some affiliations are hidden behind
	// an "Additional affiliations" button, this solves that
	affilitationSection := doc.Find("div#affiliation-section").First()

	// For each <li> in this <div> find <sup> and <span> and assign to map
	affilitationSection.Find("li").Each(func(i int, s *goquery.Selection) {
		link := strings.Trim(s.Find("sup").Text(), " ")
		affiliation := s.Find("span").Text()

		toReturn[link] = affiliation
	})

	return toReturn
}
