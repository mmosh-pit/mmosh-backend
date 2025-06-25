package ai

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
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

	if err != nil {
		log.Printf("Could not create request: %v\n", err)

		return nil, err
	}

	res, err := client.Do(request)

	if err != nil {
		log.Printf("Got error performing OpenAI request for token: %v\n", err)

		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	return body, err
}
