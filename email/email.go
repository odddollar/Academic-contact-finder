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
	address := "example@example.com"
	subject := "TestEmail"
	body := "Test Email Body"
	mailToURL := fmt.Sprintf("mailto:%s?subject=%s&body=%s", address, subject, body)
	exec.Command("powershell", "-Command", fmt.Sprintf("Start-Process '%s'", mailToURL)).Run()
}

// Send email to address of all returned results
func EmailAll() {
	SendEmail(global.A.Preferences().String("Default_email"), global.AllFoundContacts)
}
