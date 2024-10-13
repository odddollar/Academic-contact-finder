package background

import (
	"strings"

	"github.com/odddollar/CITS3200-Project/global"
	"github.com/odddollar/CITS3200-Project/widgets"
)

// Filter results based on selected option
func Filter(option string) {
	// Remove "only" from option string
	option = strings.ReplaceAll(option, " only", "")

	// Iterate through returned results and update UI
	global.Ui.Output.RemoveAll()
	for _, i := range global.AllFoundContacts {
		// Only add results based on option
		if i.Source == option {
			global.Ui.Output.Add(widgets.NewFoundContact(i))
		} else if option == "Both" {
			global.Ui.Output.Add(widgets.NewFoundContact(i))
		}
	}
}
