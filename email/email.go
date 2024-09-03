package email

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/odddollar/CITS3200-Project/global"
)

// Send data to email address
func SendEmail(address string, details []global.FoundContactStruct) {
	subject := "Requested Emails"
	var bodyContent strings.Builder
	for _, detail := range details {
		bodyContent.WriteString(fmt.Sprintf(
			"Name: %s\nSalutation: %s\nEmail: %s\nInstitution: %s\nURL: %s\n\n",
			detail.Name,
			detail.Salutation,
			detail.Email,
			detail.Institution,
			detail.URL.String(),
		))
	}
	encodedBodyContent := url.QueryEscape(bodyContent.String())
	encodedBodyContent = strings.ReplaceAll(encodedBodyContent, "+", "%20")
	mailToURL := fmt.Sprintf("mailto:%s?subject=%s&body=%s", address, subject, encodedBodyContent)
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
	// Show confirmation
	global.ShowSuccess("Email sent to: " + address)
}

// Send email to address of all returned results
func EmailAll() {
	SendEmail(global.A.Preferences().String("Default_email"), global.AllFoundContacts)
}
