package background

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/odddollar/CITS3200-Project/global"
	"github.com/odddollar/CITS3200-Project/widgets"
)

// Initiate api requesting and scraping, then update results
func Run() {
	// Ensure API key is present and valid
	// Run here again, as if cancel clicked initially then still no api key
	// This will appear until a valid key is entered every time run is clicked
	// and will not progress running any futher
	if !PresentAPIKey() || !ValidAPIKey() {
		UpdateAPIKey()
		return
	}

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

// Perform actual requesting and scraping, returning a list of found contacts
func request(firstName, lastName, institution string) []global.FoundContactStruct {
	// Create a new Chrome browser context with options to disable headless mode
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),        // Disable headless mode
		chromedp.Flag("disable-gpu", false),     // Enable GPU usage
		chromedp.Flag("start-maximized", false), // Start maximized
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// URL you want to visit
	urls := []string{
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-85079320615&origin=resultslist",
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-85070927816&origin=resultslist",
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-85037348791&origin=resultslist",
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-85119400640&origin=resultslist",
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-84922537253&origin=resultslist",
	}

	desiredName := "chris mcdonald"
	results := []global.FoundContactStruct{}

	// Iterate through urls
	for _, i := range urls {
		fmt.Println("Source: ", i)
		r := scrapeScopus(i, ctx)

		for _, j := range r {
			if fuzzy.MatchFold(j.Name, desiredName) {
				fmt.Printf("Match %s\n", j)
			} else {
				fmt.Printf("No match %s\n", j)
			}
		}

		results = append(results, r...)
	}

	return results
}

func scrapeScopus(u string, ctx context.Context) []global.FoundContactStruct {
	// Create a variable to store the page's HTML
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

	// Parse the HTML with goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		global.ShowError(err)
	}

	var toReturn []global.FoundContactStruct

	// Find all <li> elements
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		// Find all <a> elements with href attributes
		s.Find("a[href]").Each(func(j int, a *goquery.Selection) {
			href, exists := a.Attr("href")
			if exists && strings.HasPrefix(href, "mailto:") {
				// Find the <span> within the <button> to get the name
				name := s.Find("button span").Text()

				// Format results to correct structure
				up, _ := url.Parse(u)
				toReturn = append(toReturn, global.FoundContactStruct{
					Name: fmt.Sprintf("%s %s",
						strings.Split(name, ", ")[1],
						strings.Split(name, ", ")[0],
					),
					Salutation:  "Unknown",
					Email:       href[7:],
					Institution: "Unknown",
					URL:         up,
				})
			}
		})
	})

	return toReturn
}
