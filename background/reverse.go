package background

import (
	"slices"

	"github.com/odddollar/CITS3200-Project/global"
	"github.com/odddollar/CITS3200-Project/widgets"
)

func ReverseResultsOrder() {
	// Reverse order of found contacts
	slices.Reverse(global.AllFoundContacts)

	// Iterate through returned results and update UI
	global.Ui.Output.RemoveAll()
	for i := 0; i < len(global.AllFoundContacts); i++ {
		global.Ui.Output.Add(widgets.NewFoundContact(global.AllFoundContacts[i]))
	}
}
