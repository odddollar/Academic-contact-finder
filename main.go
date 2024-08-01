package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create window
	a := app.New()
	w := a.NewWindow("Academic Contact Finder")

	// Create title widget
	title := canvas.NewText("Academic Contact Finder", color.Black)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 20

	// Create search entry widgets
	firstNameLabel := widget.NewLabel("First Name:")
	firstName := widget.NewEntry()
	lastNameLabel := widget.NewLabel("Last Name:")
	lastName := widget.NewEntry()
	institutionLabel := widget.NewLabel("Institution:")
	institution := widget.NewEntry()

	// Create output box
	// TEMPORARY TEXT BOX
	// ACUTAL OUTPUT WILL HAVE MORE DATA
	output := widget.NewMultiLineEntry()

	// Create search button
	searchButton := widget.NewButton("Search", func() {})
	searchButton.Importance = widget.HighImportance

	// Create window layout
	layout := container.NewBorder(
		container.NewVBox(
			title,
			container.NewBorder(
				nil,
				nil,
				firstNameLabel,
				nil,
				firstName,
			),
			container.NewBorder(
				nil,
				nil,
				lastNameLabel,
				nil,
				lastName,
			),
			container.NewBorder(
				nil,
				nil,
				institutionLabel,
				nil,
				institution,
			),
			searchButton,
		),
		nil,
		nil,
		nil,
		output,
	)
	w.SetContent(layout)

	// Show window and run app
	w.Resize(fyne.NewSize(1024, 512))
	w.Show()
	a.Run()
}
