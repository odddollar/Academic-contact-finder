package email

import (
	"fmt"

	"github.com/odddollar/CITS3200-Project/global"
)

// Send data to email address
func SendEmail(address string, details []global.FoundContactStruct) {
	// TEMPORARY ACTION UNTIL EMAILING IS IMPLEMENTED
	fmt.Printf("Sending email to \"%s\", with data \"%v\"\n", address, details)
}

// Send email to address of all returned results
func EmailAll() {}
