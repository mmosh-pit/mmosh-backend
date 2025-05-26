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

	log.Printf("Sending request to URL: %s\n", url)

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

			log.Printf("Got info to send to client: %s\n", line)

			if len(line) > 0 {
				callbackChan <- string(line)
			}

			callbackChan <- "____break____"
			break
		}

		callbackChan <- string(line)
	}

	///

	// scanner := bufio.NewScanner(resp.Body)
	// for scanner.Scan() {
	// 	line := scanner.Text() // Or scanner.Bytes()
	// 	fmt.Printf("Received: %s\n", line)
	// 	callbackChan <- line
	//
	// 	// Here you would process the received line (e.g., parse JSON, handle event data)
	// 	// Example: if line is "event: message", "data: { ... }" for SSE
	// }
	//
	// if err := scanner.Err(); err != nil {
	// 	// An error occurred during scanning, other than io.EOF
	// 	// This could be a network error or the connection being closed unexpectedly.
	// 	if err == io.EOF {
	// 		log.Println("Stream closed by server (EOF).")
	// 		callbackChan <- "____break____"
	// 	} else {
	// 		log.Printf("Error reading stream: %v", err)
	// 		callbackChan <- "____break____"
	// 	}
	// } else {
	// 	// scanner.Scan() returned false and scanner.Err() is nil,
	// 	// which usually means io.EOF was encountered cleanly.
	// 	log.Println("Stream finished.")
	// 	callbackChan <- "____break____"
	// }
}
