package chat

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

type IncomingMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func FetchAIResponse(agentId, botId, text, systemPrompt, authToken, chatId string, namespaces []string, callbackChan chan string) {
	parsedSystemPrompt := fmt.Sprintf("%sagentId: %s, authorization: %s", systemPrompt, agentId, authToken)

	data := map[string]any{
		"query":        text,
		"namespaces":   namespaces,
		"instructions": parsedSystemPrompt,
		"agentId":      agentId,
		"bot_id":       botId,
		"aiModel":      "gpt-5.1",
	}

	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error trying to encode following payload: %v\n, %v\n", data, err)
		close(callbackChan)
		return
	}

	baseUrl := config.GetAIApiUrl()
	url := fmt.Sprintf("%s", baseUrl)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(encoded))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "text/event-stream")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	if err != nil {
		log.Printf("Error in request: %v\n", err)
		close(callbackChan)
		return
	}

	client := http.Client{
		Timeout: time.Second * 500,
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error in request: %v\n", err)
		close(callbackChan)
		return
	}

	reader := bufio.NewReader(resp.Body)

	defer resp.Body.Close()
	log.Println("Going to receive request...")

	for {
		line, err := reader.ReadBytes('\n')

		if err == nil {
			if len(line) > 0 {
				if strings.Contains(string(line), "data") {
					var message IncomingMessage

					err = json.Unmarshal(line[6:], &message)

					if err != nil {
						log.Printf("Could not parse incoming message: %v\n", err)
						continue
					}

					if message.Type == "chunk" || message.Type == "content" {
						callbackChan <- message.Content
					}
				}
			} else {
				callbackChan <- "____break____"
				break
			}

		} else {
			callbackChan <- "____break____"
			break
		}
	}
}
