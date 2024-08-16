package background

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Create new dialog with loading bar
func infiniteLoad() *dialog.CustomDialog {
	d := dialog.NewCustomWithoutButtons(
		"Loading...",
		widget.NewProgressBarInfinite(),
		global.W,
	)
	d.Resize(fyne.NewSize(300, 0))

	return d
}
