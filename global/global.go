package global

import (
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

// Struct containing data from found contacts.
// An array of these are returned by the web-scraper/api accessor
type FoundContactStruct struct {
	Name        string
	Salutation  string
	Email       string
	Institution string
	URL         *url.URL
}

// Hold global state of found contacts.
var AllFoundContacts []FoundContactStruct
