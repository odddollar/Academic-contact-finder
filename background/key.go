package background

import (
	"io"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Checks if the API is present
func PresentAPIKey() bool {
	key := global.A.Preferences().String("API_key")

	return key != ""
}

// Makes test request against API with current key. If API returns
// error then key isn't valid
func ValidAPIKey() bool {
	key := global.A.Preferences().String("API_key")
	// we need an institution token to make this work outside of the institution network
	url := "https://api.elsevier.com/content/author/author_id/57169566400?apiKey=57169566400?apiKey=" + key

	//make the http request
	resp, err := http.Get(url)
	//checks if there was an error in making the request
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Check if the HTTP response status is OK
	if resp.StatusCode != http.StatusOK {
		return false
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	bodyStr := string(body)

	// Check for specific XML error tags
	return !(strings.Contains(bodyStr, "<statusCode>AUTHENTICATION_ERROR</statusCode>") && strings.Contains(bodyStr, "<statusText>Invalid API Key</statusText>"))
}

// Opens dialog for entering new API key
func UpdateAPIKey() {
	// Get current key (will return "" if there isn't one)
	key := global.A.Preferences().String("API_key")

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
				global.A.Preferences().SetString("API_key", entry.Text)
			}
		},
		global.W,
	)
	d.Resize(fyne.NewSize(420, 0))
	d.Show()
}
