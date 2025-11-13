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

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
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

func formatChatHistory(messages []chatDomain.Message) []map[string]any {
	var formattedHistory []map[string]any
	for _, msg := range messages {
		role := "assistant"
		if msg.Type == "user" {
			role = "user"
		}

		formattedMsg := map[string]any{
			"role":      role,
			"content":   msg.Content,
			"timestamp": msg.CreatedAt,
		}
		formattedHistory = append(formattedHistory, formattedMsg)
	}

	return formattedHistory
}

func FetchAIResponse(agentId *primitive.ObjectID, botId, text, systemPrompt, authToken string, chatId *primitive.ObjectID, namespaces []string, callbackChan chan string) {

	messages := GetChatLastMessages(chatId)

	data := map[string]any{
		"query":        text,
		"namespaces":   namespaces,
		"instructions": systemPrompt,
		"chatHistory":  formatChatHistory(messages),
		"agentId":      agentId.Hex(),
		"bot_id":       botId,
	}

	log.Printf("Sending request: %v\n", data)

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

		if strings.Contains(string(line), "data") {
			var message IncomingMessage

			err = json.Unmarshal(line[6:], &message)

			if err != nil {
				log.Printf("Could not parse incoming message: %v\n", err)
				continue
			}

			if message.Type == "content" {
				callbackChan <- message.Content
			}
		}
	}
}
