package global

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

var A fyne.App
var W fyne.Window

var Ui struct {
	Title            *canvas.Image
	FirstNameLabel   *widget.Label
	FirstName        *widget.Entry
	LastNameLabel    *widget.Label
	LastName         *widget.Entry
	InstitutionLabel *widget.Label
	Institution      *widget.Entry
	Search           *widget.Button
	NumResults       *canvas.Text
	Output           *fyne.Container
}

type FoundContactStruct struct {
	Name        string
	Salutation  string
	Email       string
	Institution string
}
