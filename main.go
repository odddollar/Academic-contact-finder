package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create window
	a := app.New()
	w := a.NewWindow("Academic Contact Finder")

	// Create title widget
	title := canvas.NewText("Academic Contact Finder", color.NRGBA{86, 86, 86, 255})
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

	// Create search button
	searchButton := widget.NewButton("Search", func() {})
	searchButton.Importance = widget.HighImportance

	// Create results found label
	numResults := canvas.NewText("Found 2 results", color.NRGBA{86, 86, 86, 255})
	numResults.Alignment = fyne.TextAlignTrailing
	numResults.TextSize = 12

	// Create output box
	// TEMPORARY OUTPUT FORMAT
	output1 := widget.NewCard(
		"First Name, Last Name",
		"Example Institution",
		container.NewHBox(
			widget.NewLabel("email@example.com"),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {}),
			widget.NewButtonWithIcon("Forward", theme.MailForwardIcon(), func() {}),
			widget.NewButtonWithIcon("Website", theme.SearchIcon(), func() {}),
		),
	)
	output2 := widget.NewCard(
		"First Name, Last Name",
		"Example Institution",
		container.NewHBox(
			widget.NewLabel("email@example.com"),
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {}),
			widget.NewButtonWithIcon("Forward", theme.MailForwardIcon(), func() {}),
			widget.NewButtonWithIcon("Website", theme.SearchIcon(), func() {}),
		),
	)

	// Create window layout
	layout := container.NewVBox(
		title,
		container.NewBorder(
			nil,
			nil,
			container.NewVBox(
				firstNameLabel,
				lastNameLabel,
				institutionLabel,
			),
			nil,
			container.NewVBox(
				firstName,
				lastName,
				institution,
			),
		),
		searchButton,
		numResults,
		output1,
		output2,
	)
	w.SetContent(layout)

	// Show window and run app
	w.Resize(fyne.NewSize(1024, 0))
	w.Show()
	a.Run()
}
