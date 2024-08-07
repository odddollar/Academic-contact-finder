package widgets

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Custom widget for copying data to clipboard from entry box
type Copy struct {
	widget.Button
}

// Create Copy widget with widget to copy from
func NewCopy(w *widget.Entry) *Copy {
	copy := &Copy{}
	copy.ExtendBaseWidget(copy)

	copy.Icon = theme.ContentCopyIcon()
	copy.SetText("")
	copy.OnTapped = func() {
		// Get text from provided entry
		t := w.Text
		global.W.Clipboard().SetContent(t)
	}

	return copy
}
