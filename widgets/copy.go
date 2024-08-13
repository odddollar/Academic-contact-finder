package widgets

import (
	"fmt"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Custom widget for copying data to clipboard from entry box
type Copy struct {
	widget.Button
}

// Create Copy widget with widget to copy from
func NewCopy(s global.FoundContactStruct) *Copy {
	copy := &Copy{}
	copy.ExtendBaseWidget(copy)

	copy.Icon = theme.ContentCopyIcon()
	copy.SetText("")
	copy.OnTapped = func() {
		// Format struct to text string
		str := fmt.Sprintf("%s, %s\n%s\n%s", s.Name, s.Salutation, s.Email, s.Institution)

		// Copied formatted text to clipboard
		global.W.Clipboard().SetContent(str)
	}

	return copy
}
