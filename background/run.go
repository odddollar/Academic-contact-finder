package background

import (
	"errors"
	"fmt"
	"strings"

	"github.com/odddollar/CITS3200-Project/global"
	"github.com/odddollar/CITS3200-Project/widgets"
)

const (
	scopusApiUrl = "https://api.elsevier.com/content/search/scopus"
	googleApiUrl = "https://www.googleapis.com/customsearch/v1?"
)

// Initiate api requesting and scraping, then update results
func Run() {
	// Ensure API key is present
	// Run here again, as if cancel clicked initially then still no api key.
	// This will appear until a valid key is entered every time run is clicked
	// and will not progress running any futher
	if !PresentGoogleAPIKey() {
		UpdateGoogleAPIKey()
		return
	}

	// Get data from entry boxes
	firstName := strings.ToLower(strings.Trim(global.Ui.FirstName.Text, " "))
	lastName := strings.ToLower(strings.Trim(global.Ui.LastName.Text, " "))
	institution := strings.ToLower(strings.Trim(global.Ui.Institution.Text, " "))

	// Ensure that at least last name entered
	if lastName == "" {
		global.ShowError(errors.New("Please enter a last name"))
		return
	}

	// Create and show loading bar
	loading := infiniteLoad()
	loading.Show()

	// Clear found contact array
	global.AllFoundContacts = nil

	// Make requests and get results
	requestScopus(firstName, lastName, institution)
	requestGoogle(firstName, lastName, institution)

	// Hide loading bar
	loading.Hide()

	// Enable email all and reverse buttons if results found
	if len(global.AllFoundContacts) > 0 {
		global.Ui.EmailAll.Enable()
		global.Ui.ReverseOrder.Enable()
		global.Ui.Filter.Enable()
	}

	// Update number of results found
	global.Ui.NumResults.Text = fmt.Sprintf("Found %d results", len(global.AllFoundContacts))
	global.Ui.NumResults.Refresh()

	// Iterate through returned results and update UI
	global.Ui.Output.RemoveAll()
	for _, i := range global.AllFoundContacts {
		global.Ui.Output.Add(widgets.NewFoundContact(i))
	}
}
