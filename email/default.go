package email

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Check if a default email is present
func DefaultEmailPresent() bool {
	return global.A.Preferences().String("Default_email") != ""
}

// Change default email address
func ChangeDefaultEmail() {
	// Get current address (will return "" if there isn't one)
	address := global.A.Preferences().String("Default_email")

	// Set initial text to existing address
	entry := widget.NewEntry()
	entry.SetText(address)

	// Create form items
	options := []*widget.FormItem{
		{Text: "Email", Widget: entry, HintText: "Default address to send found contacts to"},
	}

	// Show window and update default email if "Save" selected
	d := dialog.NewForm(
		"Set default email",
		"Save",
		"Cancel",
		options,
		func(b bool) {
			if b {
				global.A.Preferences().SetString("Default_email", entry.Text)

				// Update output box to refresh emails in email me entries
				global.Ui.Output.Refresh()
			}
		},
		global.W,
	)
	d.Resize(fyne.NewSize(420, 0))
	d.Show()
}
