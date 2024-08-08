package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Custom widget for displaying box to send details to
type EmailMe struct {
	widget.BaseWidget
	details global.FoundContactStruct
	email   binding.String
}

// Create EmailMe widget with found contact details
func NewEmailMe(d global.FoundContactStruct) *EmailMe {
	// Create new string binding
	e := binding.NewString()

	email := &EmailMe{
		details: d,
		email:   e,
	}
	email.ExtendBaseWidget(email)

	return email
}

// Send data to email address
func (em *EmailMe) sendEmail() {
	// Temporary action until emailing is implemented
	address, _ := em.email.Get()
	fmt.Printf("Sending email to \"%s\", with data \"%v\"\n", address, em.details)
}

// Returns new renderer for EmailMe
func (em *EmailMe) CreateRenderer() fyne.WidgetRenderer {
	l := canvas.NewText("Email me these details:", theme.Color(theme.ColorNameForeground))
	l.TextStyle.Bold = true

	e := widget.NewEntryWithData(em.email)
	e.SetPlaceHolder("Type email address here")

	s := widget.NewButton("Send", em.sendEmail)
	s.Importance = widget.HighImportance

	return &emailMeRenderer{
		label: l,
		entry: e,
		send:  s,
	}
}

// Renderer for EmailMe widget
type emailMeRenderer struct {
	label *canvas.Text
	entry *widget.Entry
	send  *widget.Button
}

// Returns minimum size of EmailMe widget
func (r *emailMeRenderer) MinSize() fyne.Size {
	return container.NewHBox(
		r.label,
		r.entry,
		r.send,
	).MinSize()
}

func (r *emailMeRenderer) Layout(size fyne.Size) {
	// Calculate padding
	padding := theme.Padding()

	// Calculate width of entry box to make it expand
	// Find the minimum of two values to allow it to expand up to a certain point
	entrySize := fyne.NewSize(
		min(
			size.Width-r.label.MinSize().Width-r.send.MinSize().Width-4*padding,
			400,
		),
		size.Height,
	)

	// Move and resize label
	r.label.Move(fyne.NewPos(0, (size.Height-r.label.MinSize().Height)/2))
	r.label.Resize(r.label.MinSize())

	// Move and resize entry
	r.entry.Move(fyne.NewPos(r.label.MinSize().Width+padding, 0))
	r.entry.Resize(entrySize)

	// Move and resize send
	r.send.Move(fyne.NewPos(r.label.MinSize().Width+entrySize.Width+2*padding, 0))
	r.send.Resize(r.send.MinSize())
}

// Refreshes elements within widget
func (r *emailMeRenderer) Refresh() {
	r.label.Refresh()
	r.entry.Refresh()
	r.send.Refresh()
}

// Returns child elements of EmailMe
func (r *emailMeRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		r.label,
		r.entry,
		r.send,
	}
}

// Does nothing as EmailMe doesn't hold any resources
func (r *emailMeRenderer) Destroy() {}
