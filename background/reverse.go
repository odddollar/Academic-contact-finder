package background

import (
	"slices"

	"github.com/odddollar/CITS3200-Project/global"
)

func ReverseResultsOrder() {
	// Reverse order of found contacts
	slices.Reverse(global.AllFoundContacts)

	// Display based on filter option
	Filter(global.Ui.Filter.Selected)
}
