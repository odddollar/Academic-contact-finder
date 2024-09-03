package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"context"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type result struct {
	first string
	last  string
	email string
}

func main() {
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
	url := []string{
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-85079320615&origin=resultslist",
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-85070927816&origin=resultslist",
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-85037348791&origin=resultslist",
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-85119400640&origin=resultslist",
		"https://www.scopus.com/record/display.uri?eid=2-s2.0-84922537253&origin=resultslist",
	}

	firstName := "chris"
	lastName := "mcdonald"

	// Iterate through urls
	for _, i := range url {
		fmt.Println("Source: ", i)
		results := findEmails(i, ctx)

		for _, j := range results {
			if fuzzy.MatchFold(j.first, firstName) && fuzzy.MatchFold(j.last, lastName) {
				fmt.Printf("Match %s\n", j)
			} else {
				fmt.Printf("No match %s\n", j)
			}
		}

		fmt.Println()
	}
}

func findEmails(url string, ctx context.Context) []result {
	// Create a variable to store the page's HTML
	var htmlContent string

	// Visit the webpage and get the HTML content
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(1*time.Second), // Adding sleep for reliability
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the HTML with goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	var toReturn []result

	// Find all <li> elements
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		// Find all <a> elements with href attributes
		s.Find("a[href]").Each(func(j int, a *goquery.Selection) {
			href, exists := a.Attr("href")
			if exists && strings.HasPrefix(href, "mailto:") {
				// Find the <span> within the <button> to get the name
				name := s.Find("button span").Text()
				toReturn = append(toReturn, result{
					strings.Split(name, ", ")[1],
					strings.Split(name, ", ")[0],
					href[7:],
				})
			}
		})
	})

	return toReturn
}
