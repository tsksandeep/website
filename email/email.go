package email

import (
	"crypto/tls"
	b64 "encoding/base64"
	"os"

	log "github.com/sirupsen/logrus"
	gomail "gopkg.in/mail.v2"
)

const (
	smtpServer = "smtp.gmail.com"
	smtpPort   = 587

	fromEmail = "sandeepdelrio@gmail.com"
	toEmail   = "tsksandeep11@gmail.com"
	subject   = "Contact - Website"
)

func getEmailPassword() (string, error) {
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	log.Info(emailPassword)
	emailPasswordBytes, err := b64.StdEncoding.DecodeString(emailPassword)
	if err != nil {
		return "", err
	}
	return string(emailPasswordBytes), nil
}

//SendEmail send email from sandeepdelrio@gmail.com to tsksandeep11@gmail.com
func SendEmail(message string) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", fromEmail)

	// Set E-Mail receivers
	m.SetHeader("To", toEmail)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", message)

	emailPass, err := getEmailPassword()
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(emailPass)

	// Settings for SMTP server
	d := gomail.NewDialer(smtpServer, smtpPort, fromEmail, emailPass)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	return d.DialAndSend(m)
}
