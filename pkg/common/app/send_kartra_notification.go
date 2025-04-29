package common

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/sonh/qs"
)

type LeadType struct {
	FirstName string `qs:"first_name"`
	Email     string `qs:"email"`
}

type ActionType struct {
	Cmd     string `qs:"actions,cmd"`
	TagName string `qs:"actions,tag_name"`
}

type KartraRequest struct {
	AppId       string       `qs:"app_id"`
	ApiKey      string       `qs:"api_key"`
	ApiPassword string       `qs:"api_password"`
	Lead        LeadType     `qs:"lead"`
	Actions     []ActionType `qs:"actions"`
}

func SendKartraNotification(tag, name, email string) {

	appId, apiKey, apiPassword, baseUrl := config.GetKartraValues()

	d := KartraRequest{
		AppId:       appId,
		ApiKey:      apiKey,
		ApiPassword: apiPassword,
		Lead: LeadType{
			FirstName: name,
			Email:     email,
		},
	}

	encoder := qs.NewEncoder()

	values, err := encoder.Values(d)

	if err != nil {
		log.Printf("Got error trying to encode URL values to send ot kartra: %v\n", err)
		return
	}

	client := &http.Client{}

	urlVal := url.Values(values)

	urlVal.Add("actions[0][cmd]", "assign_tag")
	urlVal.Add("actions[0][tag_name]", tag)

	log.Printf("Sending data over request: %s\n", strings.NewReader(urlVal.Encode()))

	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(urlVal.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Printf("Error trying to create kartra api request: %v\n", err)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error trying to send kartra request: %v\n", err)
		return
	}

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Printf("Got error on kartra api response: %v\n", string(body))
		return
	}

	log.Printf("Kartra api request details: %v\n", string(body))
}
