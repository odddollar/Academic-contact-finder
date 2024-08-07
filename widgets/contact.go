package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Custom widget for displaying found contact information
type FoundContact struct {
	widget.BaseWidget
	name        string
	email       string
	institution string
	salutation  string
}

// Create FoundContact widget with data
func NewFoundContact(s global.FoundContactStruct) *FoundContact {
	contact := &FoundContact{
		name:        s.Name,
		email:       s.Email,
		institution: s.Institution,
		salutation:  s.Salutation,
	}
	contact.ExtendBaseWidget(contact)

	return contact
}

// Returns new renderer for FoundContact
func (fc *FoundContact) CreateRenderer() fyne.WidgetRenderer {
	return &foundContactRenderer{
		contact:     fc,
		background:  canvas.NewRectangle(color.NRGBA{128, 128, 0, 255}),
		name:        canvas.NewText(fc.name, color.Black),
		email:       canvas.NewText(fc.email, color.Black),
		institution: canvas.NewText(fc.institution, color.Black),
		salutation:  canvas.NewText(fc.salutation, color.Black),
	}
}

// Renderer for FoundContact widget
type foundContactRenderer struct {
	contact     *FoundContact
	background  *canvas.Rectangle
	name        *canvas.Text
	email       *canvas.Text
	institution *canvas.Text
	salutation  *canvas.Text
}

// Returns minimum size of FoundContact widget
func (r *foundContactRenderer) MinSize() fyne.Size {
	return container.NewVBox(
		r.name,
		r.email,
		r.institution,
		r.salutation,
	).MinSize()
}

// Lays out data and resizes FoundContact widget to fit available space
func (r *foundContactRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
}

// Refreshes elements within widget
func (r *foundContactRenderer) Refresh() {
	r.background.FillColor = color.NRGBA{128, 128, 0, 255}
	r.background.Refresh()
}

// Returns child elements of FoundContact
func (r *foundContactRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		r.background,
		r.name,
		r.email,
		r.institution,
		r.salutation,
	}
}

// Does nothing as FoundContact doesn't hold any resources
func (r *foundContactRenderer) Destroy() {}
