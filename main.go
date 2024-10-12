//go:generate fyne bundle -o static.go images/Header.png
//go:generate fyne bundle -o static.go -append images/Swap.svg

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/odddollar/CITS3200-Project/background"
	"github.com/odddollar/CITS3200-Project/email"
	"github.com/odddollar/CITS3200-Project/global"
)

func main() {
	// Create window
	global.A = app.New()
	global.A.Settings().SetTheme(global.MainTheme{})
	global.W = global.A.NewWindow("Academic Contact Finder")

	// Create title widget
	global.Ui.Title = canvas.NewImageFromResource(resourceHeaderPng)
	global.Ui.Title.FillMode = canvas.ImageFillOriginal

	// Create search entry widgets
	global.Ui.FirstNameLabel = widget.NewLabel("First name:")
	global.Ui.FirstName = widget.NewEntry()
	global.Ui.LastNameLabel = widget.NewLabel("Last name:")
	global.Ui.LastName = widget.NewEntry()
	global.Ui.InstitutionLabel = widget.NewLabel("Institution:")
	global.Ui.Institution = widget.NewEntry()

	// Create search button
	global.Ui.Search = widget.NewButtonWithIcon("Search", theme.SearchIcon(), background.Run)
	global.Ui.Search.Importance = widget.HighImportance

	// Create about button
	global.Ui.About = widget.NewButtonWithIcon("", theme.InfoIcon(), aboutCallback)

	// Create reverse results order button
	global.Ui.ReverseOrder = widget.NewButtonWithIcon("", theme.NewThemedResource(resourceSwapSvg), background.ReverseResultsOrder)
	global.Ui.ReverseOrder.Disable()

	// Create scopus and google filter buttons
	global.Ui.FilterScopus = widget.NewButtonWithIcon("", theme.NewThemedResource(resourceSwapSvg), background.FilterScopus)
	global.Ui.FilterScopus.Disable()
	global.Ui.FilterGoogle = widget.NewButtonWithIcon("", theme.NewThemedResource(resourceSwapSvg), background.FilterGoogle)
	global.Ui.FilterGoogle.Disable()
	
	// Create results found label
	global.Ui.NumResults = canvas.NewText("Found 0 results", global.Grey)
	global.Ui.NumResults.Alignment = fyne.TextAlignTrailing
	global.Ui.NumResults.TextSize = theme.CaptionTextSize()

	// Create empty container that will hold output
	global.Ui.Output = container.NewVBox()

	// Create buttons for sending emails and updating default address
	global.Ui.EmailAll = widget.NewButtonWithIcon("Email all", theme.MailSendIcon(), email.EmailAll)
	global.Ui.EmailAll.Importance = widget.HighImportance
	global.Ui.EmailAll.Disable()
	global.Ui.ChangeDefaultEmail = widget.NewButtonWithIcon("Set default email", theme.MailComposeIcon(), email.ChangeDefaultEmail)

	// Create window layout
	layout := container.NewBorder(
		container.NewVBox(
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
			container.NewBorder(
				nil,
				nil,
				nil,
				global.Ui.About,
				global.Ui.Search,
			),
			widget.NewSeparator(),
			container.NewHBox(
				global.Ui.ReverseOrder,
				global.Ui.FilterScopus,
				global.Ui.FilterGoogle,
				layout.NewSpacer(),
				global.Ui.NumResults,
			),
			widget.NewSeparator(),
		),
		container.NewBorder(
			nil,
			nil,
			nil,
			global.Ui.ChangeDefaultEmail,
			global.Ui.EmailAll,
		),
		nil,
		nil,
		container.NewScroll(global.Ui.Output),
	)
	global.W.SetContent(layout)

	// Ask if want to set default
	if !email.DefaultEmailPresent() {
		email.ChangeDefaultEmail()
	}

	// Check if google key present
	if !background.PresentGoogleAPIKey() {
		background.UpdateGoogleAPIKey()
	}

	// Check if scopus key present and valid
	if !background.PresentScopusAPIKey() || !background.ValidScopusAPIKey() {
		background.UpdateScopusAPIKey()
	}

	// Show window and run app
	global.W.Resize(fyne.NewSize(1024, 0))
	global.W.Show()
	global.A.Run()
}
