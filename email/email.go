package email

import (
	"crypto/tls"
	b64 "encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = "465"
	subject  = "Contact - Website"
)

var (
	fromEmail = mail.Address{
		Name:    "",
		Address: "sandeepdelrio@gmail.com",
	}
	toEmail = mail.Address{
		Name:    "",
		Address: "tsksandeep11@gmail.com",
	}
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
func SendEmail(body string) error {
	smtpServer := smtpServer{host: smtpHost, port: smtpPort}

	password, err := getEmailPassword()
	if err != nil {
		return err
	}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = fromEmail.String()
	headers["To"] = toEmail.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

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

	auth := smtp.PlainAuth("", fromEmail.Address, password, smtpServer.host)

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(fromEmail.Address); err != nil {
		return err
	}

	if err = c.Rcpt(toEmail.Address); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = c.Quit()
	if err != nil {
		return err
	}

	return nil
}
