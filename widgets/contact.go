package widgets

import (
	"fmt"

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
	n := canvas.NewText(
		fmt.Sprintf(
			"%s, %s",
			fc.FoundContactStruct.Name,
			fc.FoundContactStruct.Salutation,
		),
		theme.Color(theme.ColorNameForeground),
	)
	n.TextSize = theme.TextSubHeadingSize()
	n.TextStyle.Bold = true

	e := canvas.NewText(fc.FoundContactStruct.Email, theme.Color(theme.ColorNameForeground))
	e.TextSize = theme.TextSize()

	i := canvas.NewText(fc.FoundContactStruct.Institution, theme.Color(theme.ColorNameForeground))
	i.TextSize = theme.TextSize()

	se := NewEmailMe(fc.FoundContactStruct)

	c := NewCopy(fc.FoundContactStruct)

	return &foundContactRenderer{
		background:  canvas.NewRectangle(theme.Color(theme.ColorNameWarning)),
		name:        n,
		email:       e,
		institution: i,
		sendEmail:   se,
		copy:        c,
	}
}

// Renderer for FoundContact widget
type foundContactRenderer struct {
	background  *canvas.Rectangle
	name        *canvas.Text
	email       *canvas.Text
	institution *canvas.Text
	sendEmail   *EmailMe
	copy        *Copy
}

// Returns minimum size of FoundContact widget
func (r *foundContactRenderer) MinSize() fyne.Size {
	return container.NewVBox(
		container.NewHBox(
			r.name,
			r.copy,
		),
		r.email,
		r.institution,
		r.sendEmail,
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

	// Move copy button
	r.copy.Move(fyne.NewPos(size.Width-r.copy.MinSize().Width-padding, padding))
	r.copy.Resize(r.copy.MinSize())
}

// Refreshes elements within widget
func (r *foundContactRenderer) Refresh() {
	r.background.Refresh()
	r.name.Refresh()
	r.email.Refresh()
	r.institution.Refresh()
	r.sendEmail.Refresh()
	r.copy.Refresh()
}

// Returns child elements of FoundContact
func (r *foundContactRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		r.background,
		r.name,
		r.email,
		r.institution,
		r.sendEmail,
		r.copy,
	}
}

// Does nothing as FoundContact doesn't hold any resources
func (r *foundContactRenderer) Destroy() {}
