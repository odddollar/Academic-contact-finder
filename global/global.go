package global

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Main app elements
var (
	A fyne.App
	W fyne.Window
)

// Main UI elements
var Ui struct {
	Title              *canvas.Image
	FirstNameLabel     *widget.Label
	FirstName          *widget.Entry
	LastNameLabel      *widget.Label
	LastName           *widget.Entry
	InstitutionLabel   *widget.Label
	Institution        *widget.Entry
	Search             *widget.Button
	About              *widget.Button
	NumResults         *canvas.Text
	Output             *fyne.Container
	EmailAll           *widget.Button
	ChangeDefaultEmail *widget.Button
}

// Hold global state of found contacts.
var AllFoundContacts []FoundContactStruct

// Struct containing data from found contacts.
// An array of these are returned by the web-scraper/api accessor
type FoundContactStruct struct {
	Name        string
	Salutation  string
	Email       string
	Institution string
	URL         *url.URL
}

// Implement Stringer interface for FoundContactStruct
func (fcs FoundContactStruct) String() string {
	return fmt.Sprintf("%s, %s\n%s\n%s\nSource: %s", fcs.Name, fcs.Salutation, fcs.Email, fcs.Institution, fcs.URL.String())
}
