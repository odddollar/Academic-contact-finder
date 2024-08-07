package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

// Custom widget for displaying found contact information
type FoundContact struct {
	widget.BaseWidget
	name, email, institution, salutation string
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
	return &foundContactRenderer{contact: fc}
}

// Renderer for FoundContact widget
type foundContactRenderer struct {
	contact *FoundContact
}

// Returns minimum size of FoundContact widget
func (r *foundContactRenderer) MinSize() fyne.Size {
	return fyne.NewSize(10, 10)
}

// Lays out data and resizes FoundContact widget to fit available space
func (r *foundContactRenderer) Layout(size fyne.Size) {

}

// Refreshes elements within widget
func (r *foundContactRenderer) Refresh() {

}

// Returns child widgets of FoundContact
func (r *foundContactRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{}
}

// Does nothing as FoundContact doesn't hold any resources
func (r *foundContactRenderer) Destroy() {}
