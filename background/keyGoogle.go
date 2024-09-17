package background

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Checks if the Google API and search engine id is present
func PresentGoogleAPIKey() bool {
	apiKey := global.A.Preferences().String("Google_API_key")
	searchEngineID := global.A.Preferences().String("Google_search_id")

	return apiKey != "" && searchEngineID != ""
}

// Opens dialog for entering new Google API key
func UpdateGoogleAPIKey() {
	// Get current key (will return "" if there isn't one)
	apiKey := global.A.Preferences().String("Google_API_key")
	searchEngineID := global.A.Preferences().String("Google_search_id")

	// Set initial text to existing key
	apiEntry := widget.NewEntry()
	apiEntry.SetText(apiKey)
	searchEngineIDEntry := widget.NewEntry()
	searchEngineIDEntry.SetText(searchEngineID)

	// Create form items
	options := []*widget.FormItem{
		{Text: "Search engine ID", Widget: searchEngineIDEntry, HintText: "ID for custom Google search engine"},
		{Text: "API key", Widget: apiEntry, HintText: "Key for the Google API"},
	}

	// Show window and update key if "Save" selected
	d := dialog.NewForm(
		"Google API key missing, invalid, or unauthorised",
		"Save",
		"Cancel",
		options,
		func(b bool) {
			if b {
				global.A.Preferences().SetString("Google_API_key", apiEntry.Text)
				global.A.Preferences().SetString("Google_search_id", searchEngineIDEntry.Text)
			}
		},
		global.W,
	)
	d.Resize(fyne.NewSize(420, 0))
	d.Show()
}
