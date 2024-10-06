package background

import (
	"strings"
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

// func TestScopusAPIKeyEntry(t *testing.T) {
// 	// Simulate Scopus API Key entry
// 	scopusAPIEntry := widget.NewEntry()
// 	scopusAPIEntry.SetPlaceHolder("Key for Scopus API")

// 	// Simulate typing into the Scopus API key field
// 	test.Type(scopusAPIEntry, "valid-scopus-api-key")

// 	// Check if the correct key is entered
// 	if scopusAPIEntry.Text != "valid-scopus-api-key" {
// 		t.Errorf("Expected 'valid-scopus-api-key' but got '%s'", scopusAPIEntry.Text)
// 	}
// }

// func TestGoogleAPIKeyEntry(t *testing.T) {
// 	// Create form fields for Google API key and Search engine ID
// 	googleAPIEntry := widget.NewEntry()
// 	searchEngineIDEntry := widget.NewEntry()

// 	// Simulate typing the Google API key and Search Engine ID
// 	test.Type(googleAPIEntry, "valid-google-api-key")
// 	test.Type(searchEngineIDEntry, "valid-search-engine-id")

// 	// Check if the correct keys are entered
// 	if googleAPIEntry.Text != "valid-google-api-key" {
// 		t.Errorf("Expected 'valid-google-api-key' but got '%s'", googleAPIEntry.Text)
// 	}
// 	if searchEngineIDEntry.Text != "valid-search-engine-id" {
// 		t.Errorf("Expected 'valid-search-engine-id' but got '%s'", searchEngineIDEntry.Text)
// 	}
// }

func TestEmailEntry(t *testing.T) {
	// Create an email entry field
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Default address to send found contacts to")

	// Simulate typing the email
	test.Type(emailEntry, "test@example.com")

	// Check if the email is correctly entered
	if emailEntry.Text != "test@example.com" {
		t.Errorf("Expected 'test@example.com' but got '%s'", emailEntry.Text)
	}

	// Validate that the email format is correct (simple validation)
	if !strings.Contains(emailEntry.Text, "@") {
		t.Errorf("Invalid email format: %s", emailEntry.Text)
	}
}

func TestResearcherSearchForm(t *testing.T) {
	// Create the form entries
	firstNameEntry := widget.NewEntry()
	lastNameEntry := widget.NewEntry()
	institutionEntry := widget.NewEntry()

	// Simulate typing inputs
	test.Type(firstNameEntry, "Chris")
	test.Type(lastNameEntry, "Mcdonald")
	test.Type(institutionEntry, "University of Western Australia")

	// Simulate clicking the search button
	searchButton := widget.NewButton("Search", func() {
		// Simulate calling the Scopus or Google search logic here
		requestScopus(firstNameEntry.Text, lastNameEntry.Text, institutionEntry.Text)
		requestGoogle(firstNameEntry.Text, lastNameEntry.Text, institutionEntry.Text)
	})

	test.Tap(searchButton)

	// Check if the inputs are correctly set
	if firstNameEntry.Text != "Chris" {
		t.Errorf("Expected 'Chris' but got '%s'", firstNameEntry.Text)
	}
	if lastNameEntry.Text != "McDonald" {
		t.Errorf("Expected 'McDonald' but got '%s'", lastNameEntry.Text)
	}
	if institutionEntry.Text != "University of Western Australia" {
		t.Errorf("Expected 'University of Western Australia' but got '%s'", institutionEntry.Text)
	}
}
