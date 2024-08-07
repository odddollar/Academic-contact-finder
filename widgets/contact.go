package widgets

import (
	"image/color"

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
	name        string
	salutation  string
	email       string
	institution string
}

// Create FoundContact widget with data
func NewFoundContact(s global.FoundContactStruct) *FoundContact {
	contact := &FoundContact{
		name:        s.Name,
		salutation:  s.Salutation,
		email:       s.Email,
		institution: s.Institution,
	}
	contact.ExtendBaseWidget(contact)

	return contact
}

// Returns new renderer for FoundContact
func (fc *FoundContact) CreateRenderer() fyne.WidgetRenderer {
	return &foundContactRenderer{
		contact:     fc,
		background:  canvas.NewRectangle(global.LightPurple),
		name:        canvas.NewText(fc.name+", "+fc.salutation, color.Black),
		email:       canvas.NewText(fc.email, color.Black),
		institution: canvas.NewText(fc.institution, color.Black),
	}
}

// Renderer for FoundContact widget
type foundContactRenderer struct {
	contact     *FoundContact
	background  *canvas.Rectangle
	name        *canvas.Text
	email       *canvas.Text
	institution *canvas.Text
}

// Returns minimum size of FoundContact widget
func (r *foundContactRenderer) MinSize() fyne.Size {
	return container.NewVBox(
		r.name,
		r.email,
		r.institution,
	).MinSize()
}

// Lays out data and resizes FoundContact widget to fit available space
func (r *foundContactRenderer) Layout(size fyne.Size) {
	// Resize background to fill space
	r.background.Resize(size)
	r.background.CornerRadius = theme.InputRadiusSize()

	// Calculate padding
	padding := theme.Padding()
	width := size.Width - 2*padding

	// Move and resize name
	r.name.Move(fyne.NewPos(padding, padding))
	r.name.Resize(fyne.NewSize(width, r.name.MinSize().Height))

	// Move and resize email
	r.email.Move(fyne.NewPos(padding, r.name.Position().Y+r.name.Size().Height+padding))
	r.email.Resize(fyne.NewSize(width, r.email.MinSize().Height))

	// Move and resize institution
	r.institution.Move(fyne.NewPos(padding, r.email.Position().Y+r.email.Size().Height+padding))
	r.institution.Resize(fyne.NewSize(width, r.institution.MinSize().Height))
}

// Refreshes elements within widget
func (r *foundContactRenderer) Refresh() {
	r.background.Refresh()
	r.name.Refresh()
	r.email.Refresh()
	r.institution.Refresh()
}

// Returns child elements of FoundContact
func (r *foundContactRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		r.background,
		r.name,
		r.email,
		r.institution,
	}
}

// Does nothing as FoundContact doesn't hold any resources
func (r *foundContactRenderer) Destroy() {}
