package email

import (
	b64 "encoding/base64"
	"fmt"
	"net/smtp"
	"os"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"

	fromEmail = "sandeepdelrio@gmail.com"
	subject   = "Contact - Website"
)

var (
	toEmail = []string{"tsksandeep11@gmail.com"}
)

func getEmailPassword() (string, error) {
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPasswordBytes, err := b64.StdEncoding.DecodeString(emailPassword)
	if err != nil {
		return "", err
	}
	return string(emailPasswordBytes), nil
}

type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

//SendEmail send email from sandeepdelrio@gmail.com to tsksandeep11@gmail.com
func SendEmail(message string) error {
	smtpServer := smtpServer{host: smtpHost, port: smtpPort}

	password, err := getEmailPassword()
	if err != nil {
		return err
	}

	msgByte := []byte(fmt.Sprintf("To: %s\r\nSubject: Contact - Website\r\n\r\n%s", toEmail[0], message))

	auth := smtp.PlainAuth("", fromEmail, password, smtpServer.host)

	return smtp.SendMail(smtpServer.Address(), auth, fromEmail, toEmail, msgByte)
}
