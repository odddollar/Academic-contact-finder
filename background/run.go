package background

import (
	"fmt"
	"net/url"
	"time"

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
	// TEMPORARY TO SIMULATE TIME TAKEN TO PROCESS REQUEST
	time.Sleep(2 * time.Second)

	// TEMPORARY DATA
	// THIS IS WHERE THE REQUESTS TO THE API SHOULD BE INITIATED
	u, _ := url.Parse("https://example.com")
	return []global.FoundContactStruct{
		{
			Name:        "Example Example",
			Email:       "example@example.com",
			Institution: "University of example",
			Salutation:  "Dr",
			URL:         u,
		},
		{
			Name:        "Example Example2",
			Email:       "example2@example.com",
			Institution: "University of example2",
			Salutation:  "Dr",
			URL:         u,
		},
	}
}
