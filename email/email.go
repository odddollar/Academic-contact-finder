package email

import (
	"fmt"
	"os/exec"

	"github.com/odddollar/CITS3200-Project/global"
)

// Send data to email address
func SendEmail(address string, details []global.FoundContactStruct) {
	// TEMPORARY ACTION UNTIL EMAILING IS IMPLEMENTED
	fmt.Printf("Sending email to \"%s\", with data \"%v\"\n", address, details)

	// Show confirmation
	global.ShowSuccess("Email sent to: " + address)
}

func SendEmailTest() { // (address string, subject string, body string) {
	global.ShowSuccess("Start of Email Test")
	address := "example@example.com"
	subject := "TestEmail"
	body := "TestEmailBody"
	mailToURL := fmt.Sprintf("mailto:%s?subject=%s&body=%s", address, subject, body)
	exec.Command("Cmd", "/c", "start", mailToURL).Start()
	global.ShowSuccess("End of Email Test")
}

// Send email to address of all returned results
func EmailAll() {
	SendEmail(global.A.Preferences().String("Default_email"), global.AllFoundContacts)
}
