package background

import (
	"net/http"

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

// Makes test request against API with current key.
func ValidAPIKey() bool {
	key := global.A.Preferences().String("API_key")
	// we need an institution token to make this work outside of the institution network
	url := "https://api.elsevier.com/content/author/author_id/57169566400?apiKey=" + key

	//very basic api check, does not error check
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
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
