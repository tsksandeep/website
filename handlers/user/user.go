package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/website/handlers"
	"github.com/website/httputils"
	"github.com/website/sms"

	log "github.com/sirupsen/logrus"
)

var (
	ipstackAccessKey = os.Getenv("IPSTACK_ACCESS_KEY")
)

type userHanlder struct{}

// New creates a new instance of User Handler
func New() handlers.UserHandler {
	return &userHanlder{}
}

func (uh *userHanlder) PostUserDetails(w http.ResponseWriter, r *http.Request) {
	var ipAddr IPAddr

	err := json.NewDecoder(r.Body).Decode(&ipAddr)
	if err != nil {
		log.Error("error resolving relative path for sandeep resume")
		handlers.WriteHandlerError(err, http.StatusBadRequest, httputils.BadRequest, w, r)
		return
	}

	userDetails, err := getUserDetails(ipAddr.Host)
	if err != nil {
		log.Errorf("unable to send sms as the user details cannot be retrieved. %s", err)
		handlers.WriteHandlerError(err, http.StatusInternalServerError, httputils.UnexpectedError, w, r)
		return
	}

	err = sms.SendSMS(userDetails)
	if err != nil {
		log.Error(err)
		handlers.WriteHandlerError(err, http.StatusInternalServerError, httputils.UnexpectedError, w, r)
		return
	}

	err = httputils.WriteJson(200, nil, w)
	if err != nil {
		log.Error(err)
		handlers.WriteHandlerError(err, http.StatusInternalServerError, httputils.UnexpectedError, w, r)
	}

	return
}

func getUserDetails(ip string) (string, error) {
	log.Infof("%s has viewed your website", ip)

	resp, err := http.Get(fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ip, ipstackAccessKey))
	if err != nil {
		return "", err
	}

	var ipStack IPStack

	err = json.NewDecoder(resp.Body).Decode(&ipStack)
	if err != nil {
		return "", err
	}

	userDetails := fmt.Sprintf("Someone from\n%s(%s)\n%s\n%s\n%s\nhas viewed your website", ipStack.City, ipStack.Zip, ipStack.RegionName, ipStack.CountryName, ipStack.ContinentName)

	return userDetails, nil
}
