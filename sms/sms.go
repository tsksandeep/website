package sms

import (
	"net/http"
	"net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	accountSid     = os.Getenv("TWILIO_ACCOUNT_SID")
	authToken      = os.Getenv("TWILIO_AUTH_TOKEN")
	myPhoneNumber  = os.Getenv("MY_PHONE_NUMBER")
	myTwilioNumber = os.Getenv("MY_TWILIO_NUMBER")
)

// SendSMS sends SMS to my phone
func SendSMS(body string) error {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	v := url.Values{}
	v.Set("To", myPhoneNumber)
	v.Set("From", myTwilioNumber)
	v.Set("Body", body)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, err := http.NewRequest("POST", urlStr, &rb)
	if err != nil {
		return err
	}

	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	log.Info(resp.Status)

	return nil
}
