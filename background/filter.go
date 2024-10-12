package background

import (
	"github.com/odddollar/CITS3200-Project/global"
	"github.com/odddollar/CITS3200-Project/widgets"
)

// Display Scopus results only
func FilterScopus() {
	// Iterate through returned results and update UI
	global.Ui.Output.RemoveAll()
	for i := 0; i < len(global.FoundContactsScopus); i++ {
		global.Ui.Output.Add(widgets.NewFoundContact(global.FoundContactsScopus[i]))
	}
}

// Display Google results only
func FilterGoogle() {
	// Iterate through returned results and update UI
	global.Ui.Output.RemoveAll()
	for i := 0; i < len(global.FoundContactsGoogle); i++ {
		global.Ui.Output.Add(widgets.NewFoundContact(global.FoundContactsGoogle[i]))
	}
}
