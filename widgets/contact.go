package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Custom widget for displaying found contact information
type FoundContact struct {
	widget.BaseWidget
	global.FoundContactStruct
}

// Create FoundContact widget with data
func NewFoundContact(s global.FoundContactStruct) *FoundContact {
	contact := &FoundContact{
		FoundContactStruct: s,
	}
	contact.ExtendBaseWidget(contact)

	return contact
}

// Returns new renderer for FoundContact
func (fc *FoundContact) CreateRenderer() fyne.WidgetRenderer {
	n := canvas.NewText(fc.FoundContactStruct.Name+", "+fc.FoundContactStruct.Salutation, theme.Color(theme.ColorNameForeground))
	n.TextSize = theme.TextSubHeadingSize()
	n.TextStyle.Bold = true

	e := canvas.NewText(fc.FoundContactStruct.Email, theme.Color(theme.ColorNameForeground))
	e.TextSize = theme.TextSize()

	i := canvas.NewText(fc.FoundContactStruct.Institution, theme.Color(theme.ColorNameForeground))
	i.TextSize = theme.TextSize()

	se := NewEmailMe(fc.FoundContactStruct)

	return &foundContactRenderer{
		background:  canvas.NewRectangle(theme.Color(theme.ColorNameWarning)),
		name:        n,
		email:       e,
		institution: i,
		sendEmail:   se,
	}
}

// Renderer for FoundContact widget
type foundContactRenderer struct {
	background  *canvas.Rectangle
	name        *canvas.Text
	email       *canvas.Text
	institution *canvas.Text
	sendEmail   *EmailMe
}

// Returns minimum size of FoundContact widget
func (r *foundContactRenderer) MinSize() fyne.Size {
	return container.NewVBox(
		r.name,
		r.email,
		r.institution,
		r.sendEmail,
		NewSpacer(fyne.NewSize(0, theme.Padding())),
	).MinSize()
}

// Lays out data and resizes FoundContact widget to fit available space
func (r *foundContactRenderer) Layout(size fyne.Size) {
	// Resize background to fill space
	r.background.Resize(size)
	r.background.CornerRadius = theme.InputRadiusSize()

	// Calculate padding
	padding := theme.Padding()

	// Move and resize name
	r.name.Move(fyne.NewPos(padding, padding))
	r.name.Resize(r.name.MinSize())

	// Move and resize email
	r.email.Move(fyne.NewPos(padding, r.name.Position().Y+r.name.Size().Height+padding))
	r.email.Resize(r.email.MinSize())

	// Move and resize institution
	r.institution.Move(fyne.NewPos(padding, r.email.Position().Y+r.email.Size().Height+padding))
	r.institution.Resize(r.institution.MinSize())

	// Move send email
	r.sendEmail.Move(fyne.NewPos(padding, r.institution.Position().Y+r.institution.Size().Height+padding))
	r.sendEmail.Resize(fyne.NewSize(size.Width, r.sendEmail.MinSize().Height))
}

// Refreshes elements within widget
func (r *foundContactRenderer) Refresh() {
	r.background.Refresh()
	r.name.Refresh()
	r.email.Refresh()
	r.institution.Refresh()
	r.sendEmail.Refresh()
}

// Returns child elements of FoundContact
func (r *foundContactRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		r.background,
		r.name,
		r.email,
		r.institution,
		r.sendEmail,
	}
}

// Does nothing as FoundContact doesn't hold any resources
func (r *foundContactRenderer) Destroy() {}
