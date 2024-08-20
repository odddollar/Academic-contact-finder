package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"context"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func main() {
	// Create a new Chrome browser context with options to disable headless mode
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),       // Disable headless mode
		chromedp.Flag("disable-gpu", false),    // Enable GPU usage
		chromedp.Flag("start-maximized", true), // Start maximized
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// URL you want to visit
	url := "https://www.scopus.com/record/display.uri?eid=2-s2.0-85079320615&origin=resultslist&sort=plf-f&src=s&st1=McDonald&st2=Chris&nlo=1&nlr=20&nls=count-f&sid=2a239e24d72b73730a45aca45c5e6c37&sot=anl&sdt=aut&sl=39&s=AU-ID%28%22McDonald%2c+Chris+S.%22+57169566400%29&relpos=2&citeCnt=133&searchTerm="

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

	// Find all <li> elements
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		// Find all <a> elements with href attributes
		s.Find("a[href]").Each(func(j int, a *goquery.Selection) {
			href, exists := a.Attr("href")
			if exists && strings.HasPrefix(href, "mailto:") {
				// Find the <span> within the <button> to get the name
				name := s.Find("button span").Text()
				fmt.Printf("%s: %s\n", name, href[7:])
			}
		})
	})
}
