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
	ReverseOrder       *widget.Button
	FilterScopus       *widget.Button
	FilterGoogle       *widget.Button
	NumResults         *canvas.Text
	Output             *fyne.Container
	EmailAll           *widget.Button
	ChangeDefaultEmail *widget.Button
}

// Hold global state of found contacts.
var AllFoundContacts []FoundContactStruct
var FoundContactsScopus []FoundContactStruct
var FoundContactsGoogle []FoundContactStruct

// Struct containing data from found contacts.
// An array of these are returned by the web-scraper/api accessor
type FoundContactStruct struct {
	FirstName   string
	LastName    string
	Salutation  string
	Email       string
	Institution string
	URL         *url.URL
}

// Implement Stringer interface for FoundContactStruct
func (fcs FoundContactStruct) String() string {
	// Change name formatting if salutation provided
	var nameSal string
	if fcs.Salutation == "" {
		nameSal = fmt.Sprintf(
			"%s %s",
			fcs.FirstName,
			fcs.LastName,
		)
	} else {
		nameSal = fmt.Sprintf(
			"%s %s, %s",
			fcs.FirstName,
			fcs.LastName,
			fcs.Salutation,
		)
	}

	return fmt.Sprintf("%s\n%s\n%s\nSource: %s\n\n", nameSal, fcs.Email, fcs.Institution, fcs.URL.String())
}
