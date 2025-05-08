package common

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func FetchAIResponse(username, text, systemPrompt string, namespaces []string) string {
	data := map[string]interface{}{
		"username":      username,
		"prompt":        text,
		"namespaces":    namespaces,
		"system_promtp": systemPrompt,
	}

	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error trying to encode following payload: %v\n, %v\n", data, err)
		return ""
	}

	baseUrl := config.GetAIApiUrl()

	url := fmt.Sprintf("%s/generate/", baseUrl)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(encoded))

	if err != nil {
		log.Printf("Error in request: %v\n", err)
		return ""
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error in request: %v\n", err)
		return ""
	}

	reader := bufio.NewReader(resp.Body)

	defer resp.Body.Close()

	result := ""

	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {

			if len(line) > 0 {
				result += string(line)
			}
			break
		}

		result += string(line)
	}

	return result
}
