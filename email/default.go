package email

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Change default email address
func ChangeDefaultEmail() {
	// Get current address (will return "" if there isn't one)
	address := global.A.Preferences().String("Default_email")

	// Set initial text to existing address
	entry := widget.NewEntry()
	entry.Validator = validation.NewRegexp(
		"^[a-zA-Z0-9.!#$%&'\"*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
		"Must be valid email address",
	)
	entry.SetText(address)

	// Create form items
	options := []*widget.FormItem{
		{Text: "Email", Widget: entry, HintText: "Default address to send found contacts to"},
	}

	// Show window and update default email if "Save" selected
	d := dialog.NewForm(
		"Change default email",
		"Save",
		"Cancel",
		options,
		func(b bool) {
			if b {
				global.A.Preferences().SetString("Default_email", entry.Text)
			}
		},
		global.W,
	)
	d.Resize(fyne.NewSize(420, 0))
	d.Show()
}
