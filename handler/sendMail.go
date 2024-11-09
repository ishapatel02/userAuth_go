package handler

import (
	"fmt"
	"net/smtp"
)

// SendEmail sends an email with the provided subject and body to the specified recipient(s).
func SendEmail(subject, body string, to []string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "ishapatel2021@gmail.com", "nsli rzda njww zkfx", "smtp.gmail.com")

	// Define the sender's email.
	from := "ishapatel2021@gmail.com"

	// Compose the email message.
	message := []byte("Subject: " + subject + "\n\n" + body)

	// Send the email.
	// err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)

	fmt.Println("Mail Send")
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
