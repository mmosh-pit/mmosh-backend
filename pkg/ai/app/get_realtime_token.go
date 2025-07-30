package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetRealtimeToken() (any, error) {

	client := &http.Client{}

	reqBody := map[string]string{
		"model": "gpt-4o-realtime-preview-2025-06-03",
		"voice": "verse",
	}

	encoded, err := json.Marshal(reqBody)

	if err != nil {
		log.Printf("Could not encode request body: %v\n", err)

		return nil, err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/realtime/sessions", bytes.NewBuffer(encoded))

	key := config.GetOpenAIKey()

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))

	request.Header.Add("Content-Type", "application/json")

	log.Printf("Sending headers: %v\n", request.Header.Get("Authorization"))

	if err != nil {
		log.Printf("Could not create request: %v\n", err)

		return nil, err
	}

	res, err := client.Do(request)

	if err != nil {
		log.Printf("Got error performing OpenAI request for token: %v\n", err)

		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, errors.New("unauthorized")
	}

	body, err := io.ReadAll(res.Body)

	log.Printf("response body: %v\n", string(body))
	log.Printf("Response code: %v\n", res.StatusCode)

	var result any

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}
