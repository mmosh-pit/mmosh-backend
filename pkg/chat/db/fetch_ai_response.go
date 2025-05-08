package chat

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func FetchAIResponse(username, text, systemPrompt string, namespaces []string, callbackChan chan string) {
	data := map[string]interface{}{
		"username":      username,
		"prompt":        text,
		"namespaces":    namespaces,
		"system_prompt": systemPrompt,
		"model":         "gemini-2.0-flash",
	}

	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error trying to encode following payload: %v\n, %v\n", data, err)
		close(callbackChan)
		return
	}

	baseUrl := config.GetAIApiUrl()

	url := fmt.Sprintf("%s/generate_stream/", baseUrl)

	log.Printf("Sending AI request with body: %v\n", string(encoded))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(encoded))

	if err != nil {
		log.Printf("Error in request: %v\n", err)
		close(callbackChan)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error in request: %v\n", err)
		close(callbackChan)
		return
	}

	reader := bufio.NewReader(resp.Body)

	defer resp.Body.Close()

	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {

			if len(line) > 0 {
				callbackChan <- string(line)
			}

			callbackChan <- "____break____"
			break
		}

		callbackChan <- string(line)
	}
}
