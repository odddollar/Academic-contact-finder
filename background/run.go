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
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/odddollar/CITS3200-Project/global"
	"github.com/odddollar/CITS3200-Project/widgets"
)

// Scruct created to store list of URLs found
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

// Initiate api requesting and scraping, then update results
func Run() {
	// Ensure API key is present and valid
	// Run here again, as if cancel clicked initially then still no api key
	// This will appear until a valid key is entered every time run is clicked
	// and will not progress running any futher
	if !PresentScopusAPIKey() || !ValidScopusAPIKey() {
		UpdateScopusAPIKey()
		return
	}

	// Get data from entry boxes
	firstName := strings.ToLower(global.Ui.FirstName.Text)
	lastName := strings.ToLower(global.Ui.LastName.Text)
	institution := strings.ToLower(global.Ui.Institution.Text)

	// Create and show loading bar
	loading := infiniteLoad()
	loading.Show()

	// Make request and get results
	request(firstName, lastName, institution)

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

// Perform actual requesting and scraping, returning a list of found contacts
func request(firstName, lastName, institution string) {
	// Create a new Chrome browser context with options to disable headless mode
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

	firstName = "chris"
	lastName = "mcdonald"
	institution = "university of western australia"

	// Build the request URL
	apiUrl := "https://api.elsevier.com/content/search/scopus"
	apiKey := global.A.Preferences().String("Scopus_API_key")

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
	fmt.Println(string(body))

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		global.ShowError(errors.New("bad http response"))
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

	results := []global.FoundContactStruct{}

	// Iterate through urls
	for _, i := range urls {
		// Scrape data from current url
		r := scrapeScopus(i, ctx)

		// Iterate through results from url
		for _, j := range r {
			fnd := fuzzy.LevenshteinDistance(strings.ToLower(j.FirstName), firstName)
			lnd := fuzzy.LevenshteinDistance(strings.ToLower(j.LastName), lastName)
			//fmt.Printf("%s %s: %d %d\n", j.FirstName, j.LastName, fnd, lnd)

			// Check if result qualifies as match
			if fnd <= 2 || lnd <= 2 {
				results = append(results, j)
			}
		}
	}

	global.AllFoundContacts = results
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

				// Find the <sup> to get affiliation link
				affiliationLink := s.Find("sup").Text()
				affiliation := affiliations[affiliationLink]
				if affiliation == "" {
					affiliation = "Unknown Institution"
				}

				// Format results to correct structure
				up, _ := url.Parse(u)
				toReturn = append(toReturn, global.FoundContactStruct{
					FirstName:   strings.Split(name, ", ")[1],
					LastName:    strings.Split(name, ", ")[0],
					Salutation:  "Unknown Salutation", // Salutation not provided by scopus
					Email:       href[7:],             // Remove "mailto:"
					Institution: affiliation,          // Get affiliation from map
					URL:         up,                   // Parsed url as source
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
