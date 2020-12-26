package email

import (
	"crypto/tls"
	b64 "encoding/base64"
	"fmt"
	"net/smtp"
	"os"
)

const (
	smtpHost        = "smtp.ionos.com"
	smtpPort        = "465"
	fromEmail       = "admin@tsksandeep.com"
	toEmailPersonal = "tsksandeep11@gmail.com"
	subjectPersonal = "Contact - Website"
	subjectClient   = "Thank you for contacting"
	messageClient   = "Thank you for contacting me. I will get back to you as soon as possible. Please contact me in this number (+91 9442142327) if anything urgent.\nThanks,\nSandeep Kumar"
)

func emailWriter(c *smtp.Client, fromEmail string, toEmail string, message string) error {
	if err := c.Mail(fromEmail); err != nil {
		return err
	}

	if err := c.Rcpt(toEmail); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	defer w.Close()

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	return nil
}

func getEmailPassword() (string, error) {
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPasswordBytes, err := b64.StdEncoding.DecodeString(emailPassword)
	if err != nil {
		return "", err
	}
	return string(emailPasswordBytes), nil
}

func getEmailMessage(fromEmail string, toEmail string, subject string, body string) string {
	headers := make(map[string]string)
	headers["From"] = fromEmail
	headers["To"] = toEmail
	headers["Subject"] = subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	return message
}

//SendEmail sends email to the owner's account and the user's account
func SendEmail(body Body) error {
	smtpServer := smtpServer{host: smtpHost, port: smtpPort}

	password, err := getEmailPassword()
	if err != nil {
		return err
	}

	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.Address(), tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		return err
	}

	defer c.Quit()

	auth := smtp.PlainAuth("", fromEmail, password, smtpServer.host)

	if err = c.Auth(auth); err != nil {
		return err
	}

	// Send message to personal account
	emailWriter(c, fromEmail, toEmailPersonal, getEmailMessage(fromEmail, toEmailPersonal, subjectPersonal, body.ToString()))

	msgClient := fmt.Sprintf("Hi %s,\n%s", body.GetName(), messageClient)

	// Send message to client account
	emailWriter(c, fromEmail, body.GetEmail(), getEmailMessage(fromEmail, body.GetEmail(), subjectClient, msgClient))

	return nil
}
