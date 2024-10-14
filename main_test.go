package background

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Testing the api validity as well as the buttons and interface

// Function testing the google API key validity
func isGoogleAPIKeyValid(apiKey string) bool {

	searchEngineID := "testinvalidengineID"
	searchQuery := "Chris McDonald UWA"
	// Build the query parameters
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("cx", searchEngineID)
	params.Add("q", searchQuery)
	params.Add("start", fmt.Sprintf("%d", 2))
	params.Add("num", fmt.Sprintf("%d", 3))

	// Build final url with parameters
	reqUrl := fmt.Sprintf("%s%s", googleApiUrl, params.Encode())

	// Send GET request to Google Search API
	resp, err := http.Get(reqUrl)
	if err != nil {
		global.ShowError(err)
	}
	defer resp.Body.Close()

	// Check if response is successful.
	// If not then key probably isn't valid
	if resp.StatusCode != http.StatusOK {
		// Fixes dialog disappearing immediately
		go func() {
			UpdateGoogleAPIKey()
		}()
		return false
	}
	return true
}

// Function testing the google API key functions and validity testing
func TestUpdateGoogleAPIKey_InvalidKey(t *testing.T) {
	// Create a new test app with a unique ID
	a := app.NewWithID("com.example.uniqueID")
	w := a.NewWindow("Test Window")

	// Simulate the Preferences storage
	preferences := a.Preferences()

	// Set initial values for API key and search engine ID
	preferences.SetString("Google_API_key", "existing-api-key")
	preferences.SetString("Google_search_id", "existing-search-id")

	// Function to simulate the dialog
	UpdateGoogleAPIKey := func() {
		apiKey := preferences.String("Google_API_key")
		searchEngineID := preferences.String("Google_search_id")

		// Set initial text to existing key
		apiEntry := widget.NewEntry()
		apiEntry.SetText(apiKey)
		searchEngineIDEntry := widget.NewEntry()
		searchEngineIDEntry.SetText(searchEngineID)

		// Create form items
		options := []*widget.FormItem{
			{Text: "Instructions", Widget: widget.NewLabel("Please refer to user guide")},
			{Text: "Search engine ID", Widget: searchEngineIDEntry, HintText: "ID for custom Google search engine"},
			{Text: "API key", Widget: apiEntry, HintText: "Key for Google API"},
		}

		// Create the dialog
		d := dialog.NewForm(
			"Google API key missing, invalid, or unauthorised",
			"Save",
			"Cancel",
			options,
			func(b bool) {
				if b {
					// Simulate checking the API key validity
					apikeytest := isAPIKeyValid(apiEntry.Text)
					if apikeytest != true {
						t.Errorf("API key is invalid: %v", apikeytest)
					} else {
						// Update the preferences if the key is valid
						preferences.SetString("Google_API_key", apiEntry.Text)
						preferences.SetString("Google_search_id", searchEngineIDEntry.Text)
					}
				}
			},
			w,
		)

		// Show the dialog headlessly
		d.Resize(fyne.NewSize(420, 0))
		d.Show()
	}

	// Call the function that triggers the dialog
	UpdateGoogleAPIKey()

	// Simulate entering an invalid API key
	apiEntry := widget.NewEntry()
	apiEntry.SetText("invalid-api-key")
	test.Type(apiEntry, "invalid-api-key")

	// Simulate clicking the Save button
	saveButton := widget.NewButton("Save", nil)
	test.Tap(saveButton)

	// Check that the API key was not updated due to being invalid
	if preferences.String("Google_API_key") != "existing-api-key" {
		t.Errorf("Expected 'existing-api-key', got '%s'", preferences.String("Google_API_key"))
	}
}

// Testing the Scopus api key validity as well as the buttons and interface

func isScopusAPIKeyValid(apiKey string) bool {
	key := apiKey

	// Need an institution token to make this work outside of the institution network
	url := "https://api.elsevier.com/content/author/author_id/57169566400?apiKey=" + key

	// Very basic api check
	resp, err := http.Get(url)
	if err != nil {
		global.ShowError(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// Function testing the Scopys API key functions and validity testing
func TestUpdateScopusAPIKey_InvalidKey(t *testing.T) {
	// Create a new test app with a unique ID
	a := app.NewWithID("com.example.uniqueID")
	w := a.NewWindow("Test Window")

	// Simulate the Preferences storage
	preferences := a.Preferences()

	// Set initial values for API key and search engine ID
	preferences.SetString("Scopus_API_key", "existing-api-key")

	// Function to simulate the dialog
	UpdateScopusAPIKey := func() {
		apiKey := preferences.String("Scopus_API_key")

		// Set initial text to existing key
		apiEntry := widget.NewEntry()
		apiEntry.SetText(apiKey)

		// Create form items
		options := []*widget.FormItem{
			{Text: "Instructions", Widget: widget.NewLabel("Please refer to user guide")},
			{Text: "API key", Widget: apiEntry, HintText: "Key for Scopus API"},
		}

		// Create the dialog
		d := dialog.NewForm(
			"Scopus API key missing, invalid, or unauthorised",
			"Save",
			"Cancel",
			options,
			func(b bool) {
				if b {
					// Simulate checking the API key validity
					apikeytest := isScopusAPIKeyValid(apiEntry.Text)
					if apikeytest != true {
						t.Errorf("API key is invalid: %v", apikeytest)
					} else {
						// Update the preferences if the key is valid
						preferences.SetString("Scopus_API_key", apiEntry.Text)
					}
				}
			},
			w,
		)

		// Show the dialog headlessly
		d.Resize(fyne.NewSize(420, 0))
		d.Show()
	}

	// Call the function that triggers the dialog
	UpdateScopusAPIKey()

	// Simulate entering an invalid API key
	apiEntry := widget.NewEntry()
	apiEntry.SetText("invalid-api-key")
	test.Type(apiEntry, "invalid-api-key")

	// Simulate clicking the Save button
	saveButton := widget.NewButton("Save", nil)
	test.Tap(saveButton)

	// Check that the API key was not updated due to being invalid
	if preferences.String("Scopus_API_key") != "existing-api-key" {
		t.Errorf("Expected 'existing-api-key', got '%s'", preferences.String("Scopus_API_key"))
	}
}

// Testing the logic and reliability of the Google scraping
func TestGoogleScraping(t *testing.T) {
	var firstName string
	var lastName string
	var institution string

	firstName = "chris"
	lastName = "mcdonald"
	institution = "uwa"

	fileName := "testgooglescraping.html"

	// Read the file into a byte slice
	htmlData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Convert byte slice to string
	htmlContent := string(htmlData)
	// Convert html to lowercase
	htmlContentLower := strings.ToLower(htmlContent)

	// Parse email regex
	emailRe := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find all emails and filter down to matching ones
	foundEmails := emailRe.FindAllString(htmlContentLower, -1)
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

	// Compile salutation regex and search for highest salutation
	salutationRe := regexp.MustCompile(`(?i)(\b(?:Prof\.?|Professor)\b|\b(?:Assoc(?:\.|\b)\s*(?:Prof(?:\.|\b)|Professor)|Associate\s*(?:Prof(?:\.|\b)|Professor))\b|\b(?:Asst(?:\.|\b)\s*(?:Prof(?:\.|\b)|Professor)|Assistant\s*(?:Prof(?:\.|\b)|Professor))\b|\b(?:Dr\.?|Doctor)\b)`)
	salutation := getHighestSalutation(salutationRe.FindAllString(htmlContent, -1))

	// Format results to correct structure
	up, _ := url.Parse("www.testing.com")
	result := global.FoundContactStruct{
		FirstName:   firstName,
		LastName:    lastName,
		Salutation:  salutation,
		Email:       email,
		Institution: institution,
		URL:         up,
	}

	// Define the expected email address
	expectedEmail := "Chris.McDonald@uwa.edu.au"

	// Compare the found email with the expected email
	if strings.ToLower(result.Email) != strings.ToLower(expectedEmail) {
		t.Errorf("Expected email '%s', but found '%s'", expectedEmail, result.Email)
	}

}

// Testing the logic and reliability of the Scopus scraping

func TestScopusScraping(t *testing.T) {

	// Read the file into a byte slice
	htmlData, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Convert byte slice to string
	htmlContent := string(htmlData)

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

				// If no affiliation <sup> found, then only one affiliation in map
				if affiliation == "" {
					affiliation = affiliations["a"]
				}

				// Format results to correct structure
				up, _ := url.Parse("www.testwebsite.com")
				result = append(result, global.FoundContactStruct{
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
	// Define the expected email address
	expectedEmail := "Chris.McDonald@uwa.edu.au"

	// Compare the found email with the expected email
	if strings.ToLower(result.Email) != strings.ToLower(expectedEmail) {
		t.Errorf("Expected email '%s', but found '%s'", expectedEmail, result.Email)
	}

}
