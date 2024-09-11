package main

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Show a dialog box with about information
func aboutCallback() {
	// Button to clear api key if desired
	b := widget.NewButton("Clear settings and API keys", func() {
		global.A.Preferences().SetString("Scopus_API_key", "")
		global.A.Preferences().SetString("Default_email", "")

		// Show success dialog
		global.ShowSuccess("Settings and API keys cleared")
	})

	// Separate markdown widgets for better spacing
	d := container.NewVBox(
		widget.NewRichTextFromMarkdown(fmt.Sprintf("**Version**: %s", global.A.Metadata().Version)),
		widget.NewRichTextFromMarkdown("**Created for**: CITS3200 Professional Computing, Semester 2 2024, UWA"),
		widget.NewRichTextFromMarkdown("**By**: Kyle Dunstall, Lucy Dye, Simon Eason, Ryan Kingsley, Sersang Ngedup, James Wigfield"),
		b,
	)

	// Show dialog with layout
	dialog.ShowCustom(
		"About",
		"OK",
		d,
		global.W,
	)
}
