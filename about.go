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
	// Separate markdown widgets for better spacing
	d := container.NewVBox(
		widget.NewRichTextFromMarkdown(fmt.Sprintf("**Version**: %s", global.A.Metadata().Version)),
		widget.NewRichTextFromMarkdown("**Created for**: CITS3200 Professional Computing, Semester 2 2024, UWA"),
		widget.NewRichTextFromMarkdown("**By**: Kyle Dunstall, Lucy Dye, Simon Eason, Ryan Kingsley, Sersang Ngedup, James Wigfield"),
	)

	// Show dialog with layout
	dialog.ShowCustom(
		"About",
		"OK",
		d,
		global.W,
	)
}
