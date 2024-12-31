package chat

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func FetchAIResponse(username, text string, namespaces []string, callbackChan chan string) {
	data := map[string]interface{}{
		"username":   username,
		"prompt":     text,
		"namespaces": namespaces,
	}

	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error trying to encode following payload: %v\n, %v\n", data, err)
		close(callbackChan)
		return
	}

	url := "https://mmoshapi-uodcouqmia-uc.a.run.app/generate_stream/"

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

	fmt.Println("Response: Content-length:", resp.Header.Get("Content-length"))

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
