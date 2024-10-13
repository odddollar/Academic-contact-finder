package background

import (
	"github.com/odddollar/CITS3200-Project/global"
	"github.com/odddollar/CITS3200-Project/widgets"
)

// Display Scopus results only
func FilterScopus() {
	// Iterate through returned results and update UI
	global.Ui.Output.RemoveAll()
	for _, i := range global.AllFoundContacts {
		if i.Source == "Scopus" {
			global.Ui.Output.Add(widgets.NewFoundContact(i))
		}
	}
}

// Display Google results only
func FilterGoogle() {
	// Iterate through returned results and update UI
	for _, i := range global.AllFoundContacts {
		if i.Source == "Google" {
			global.Ui.Output.Add(widgets.NewFoundContact(i))
		}
	}
}
