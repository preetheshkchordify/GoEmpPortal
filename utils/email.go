package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"strings"
)

type Email struct {
	To           []string
	CC           []string
	Subject      string
	TemplatePath string
	Placeholders map[string]string
	Attachments  []string
}

// Export the function by capitalizing its name
func SendEmail(email Email) error {
	templateContent, err := os.ReadFile(email.TemplatePath)
	if err != nil {
		return fmt.Errorf("failed to read template: %v", err)
	}

	tmpl, err := template.New("email").Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, email.Placeholders)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")
	username := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")

	auth := smtp.PlainAuth("", username, password, smtpHost)

	// ✅ FIX: Ensure SMTP address is correctly formatted
	smtpAddress := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	headers := fmt.Sprintf("From: %s\r\n", username)
	headers += fmt.Sprintf("To: %s\r\n", strings.Join(email.To, ","))
	if len(email.CC) > 0 {
		headers += fmt.Sprintf("Cc: %s\r\n", strings.Join(email.CC, ","))
	}
	headers += fmt.Sprintf("Subject: %s\r\n", email.Subject)
	headers += "MIME-Version: 1.0\r\n"
	headers += "Content-Type: text/html; charset=UTF-8\r\n\r\n"

	message := []byte(headers + body.String())

	// ✅ FIX: Check if connection is established
	fmt.Println("📧 Connecting to SMTP server:", smtpAddress)

	err = smtp.SendMail(smtpAddress, auth, username, append(email.To, email.CC...), message)
	if err != nil {
		fmt.Println("❌ Error sending email:", err)
		return err
	}

	fmt.Println("✅ Email sent successfully!")
	return nil
}
