package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/global"
)

func main() {
	// Create window
	global.A = app.New()
	global.W = global.A.NewWindow("Academic Contact Finder")

	// Create title widget
	global.Ui.Title = canvas.NewText("Academic Contact Finder", global.TextColour)
	global.Ui.Title.Alignment = fyne.TextAlignCenter
	global.Ui.Title.TextStyle.Bold = true
	global.Ui.Title.TextSize = 20

	// Create search entry widgets
	global.Ui.FirstNameLabel = widget.NewLabel("First Name:")
	global.Ui.FirstName = widget.NewEntry()
	global.Ui.LastNameLabel = widget.NewLabel("Last Name:")
	global.Ui.LastName = widget.NewEntry()
	global.Ui.InstitutionLabel = widget.NewLabel("Institution:")
	global.Ui.Institution = widget.NewEntry()

	// Create search button
	global.Ui.Search = widget.NewButtonWithIcon("Search", theme.SearchIcon(), func() {})

	// Create results found label
	global.Ui.NumResults = canvas.NewText("Found 0 results", global.TextColour)
	global.Ui.NumResults.Alignment = fyne.TextAlignTrailing
	global.Ui.NumResults.TextSize = 12

	// Create empty container that will hold output
	global.Ui.Output = container.NewVBox()

	// Create window layout
	layout := container.NewVBox(
		global.Ui.Title,
		container.NewBorder(
			nil,
			nil,
			container.NewVBox(
				global.Ui.FirstNameLabel,
				global.Ui.LastNameLabel,
				global.Ui.InstitutionLabel,
			),
			nil,
			container.NewVBox(
				global.Ui.FirstName,
				global.Ui.LastName,
				global.Ui.Institution,
			),
		),
		container.NewThemeOverride(global.Ui.Search, global.ButtonTheme{}),
		widget.NewSeparator(),
		global.Ui.NumResults,
		global.Ui.Output,
	)
	global.W.SetContent(layout)

	// Show window and run app
	global.A.Settings().SetTheme(global.MainTheme{})
	global.W.Resize(fyne.NewSize(1024, 0))
	global.W.Show()
	global.A.Run()
}
