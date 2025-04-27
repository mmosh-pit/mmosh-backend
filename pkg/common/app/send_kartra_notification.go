package common

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SendKartraNotification(tag, name, email string) {

	appId, apiKey, apiPassword, baseUrl := config.GetKartraValues()

	data := map[string]interface{}{
		"app_id":       appId,
		"api_key":      apiKey,
		"api_password": apiPassword,
		"lead": map[string]string{
			"first_name": name,
			"email":      email,
		},
		"actions": []map[string]string{
			{
				"cmd":      "assign_tag",
				"tag_name": tag,
			},
		},
	}

	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error trying to send kartra notification: %v\n", err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(encoded))

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
