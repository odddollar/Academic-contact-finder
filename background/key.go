package background

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Checks if the Scopus API is present
func PresentScopusAPIKey() bool {
	key := global.A.Preferences().String("Scopus_API_key")

	return key != ""
}

// Makes test request against Scopus API with current key. If API returns
// error then key isn't valid
func ValidScopusAPIKey() bool {
	key := global.A.Preferences().String("Scopus_API_key")

	// UPDATE HERE TO MAKE TEST REQUEST TO ENSURE API KEY WORKS
	return key != ""
}

// Opens dialog for entering new Scopus API key
func UpdateScopusAPIKey() {
	// Get current key (will return "" if there isn't one)
	key := global.A.Preferences().String("Scopus_API_key")

	// Set initial text to existing key
	entry := widget.NewEntry()
	entry.SetText(key)

	// Create form items
	options := []*widget.FormItem{
		{Text: "API key", Widget: entry, HintText: "Key for the Scopus API"},
	}

	// Show window and update key if "Save" selected
	d := dialog.NewForm(
		"API key missing or invalid",
		"Save",
		"Cancel",
		options,
		func(b bool) {
			if b {
				global.A.Preferences().SetString("Scopus_API_key", entry.Text)
			}
		},
		global.W,
	)
	d.Resize(fyne.NewSize(420, 0))
	d.Show()
}
