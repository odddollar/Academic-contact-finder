package email

import (
	"fmt"
	"os/exec"
	"runtime"

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
	switch runtime.GOOS {
	case "windows":
		// Use PowerShell to open the mailto URL on Windows
		exec.Command("powershell", "-Command", fmt.Sprintf("Start-Process '%s'", mailToURL)).Run()
	case "linux":
		// Use xdg-open to open the mailto URL on Linux
		exec.Command("xdg-open", mailToURL).Run()
	case "darwin":
		// Use open to open the mailto URL on macOS
		exec.Command("open", mailToURL).Run()
	}
}

// Send email to address of all returned results
func EmailAll() {
	SendEmail(global.A.Preferences().String("Default_email"), global.AllFoundContacts)
}
