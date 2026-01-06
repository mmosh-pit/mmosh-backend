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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// const queryData = {
//   namespaces: [selectedChat!.chatAgent!.key, "PUBLIC"],
//   query: content,
//   instructions: selectedChat!.chatAgent!.system_prompt,
//   chatHistory: chatHistory,
//   agentId: selectedChat.chatAgent!.id,
//   bot_id: selectedChat.chatAgent!.key,
// };

type IncomingMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// func formatChatHistory(messages []chatDomain.Message) []map[string]any {
// 	var formattedHistory []map[string]any
// 	for _, msg := range messages {
// 		role := "assistant"
// 		if msg.Type == "user" {
// 			role = "user"
// 		}
//
// 		formattedMsg := map[string]any{
// 			"role":      role,
// 			"content":   msg.Content,
// 			"timestamp": msg.CreatedAt,
// 		}
// 		formattedHistory = append(formattedHistory, formattedMsg)
// 	}
//
// 	return formattedHistory
// }

func FetchAIResponse(agentId *primitive.ObjectID, botId, text, systemPrompt, authToken string, chatId *primitive.ObjectID, namespaces []string, callbackChan chan string) {

	parsedSystemPrompt := fmt.Sprintf("%sagentId: %s, authorization: %s", systemPrompt, agentId.Hex(), authToken)

	data := map[string]any{
		"query":        text,
		"namespaces":   namespaces,
		"instructions": parsedSystemPrompt,
		// "chatHistory":  formatChatHistory(messages),
		"agentId": agentId.Hex(),
		"bot_id":  botId,
		"aiModel": "gpt-5.1",
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
