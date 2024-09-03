package global

import (
	"fyne.io/fyne/v2/dialog"
)

// Standard dialog to show success
func ShowSuccess(message string) {
	dialog.ShowInformation("Success", message, W)
}

// Standard dialog to show error
func ShowError(err error) {
	dialog.ShowInformation("Error", err.Error(), W)
}
