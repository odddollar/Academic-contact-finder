package global

import "fyne.io/fyne/v2/dialog"

// Standard dialog to show success
func ShowSuccess(message string) {
	// Show confirmation dialog
	dialog.ShowInformation("Success", message, W)
}
