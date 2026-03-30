package chat

import (
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
)

func GenerateAIResponse(client *PoolClient, message *chatDomain.Message) {
	log.Println("0")

	textChan := make(chan string)

	tempId := fmt.Sprintf("%d", time.Now().UnixNano())

	chat, _ := chatDb.GetChatById(message.ChatId)

	go chatDb.SaveMessage(message, chat.Agent.Key, "", client.Token)
	go chatDb.UpdateChatLastMessage(message)

	index := randRange(0, 2)

	senderId := botParticipantUsers[index].ID

	if message.AgentId != "" {
		senderId = message.AgentId
	}

	loadingMessage := chatDomain.Message{
		ID:        tempId,
		Content:   "",
		Type:      "bot",
		Sender:    senderId,
		IsLoading: true,
		AgentId:   message.AgentId,
		ChatId:    message.ChatId,
	}

	data := map[string]any{
		"event": "aiMessage",
		"data":  loadingMessage,
	}

	client.sendResponse(data)

	createdDate := time.Now()

	generatedMessage := chatDomain.Message{
		ID:        tempId,
		Content:   "",
		Type:      "bot",
		Sender:    senderId,
		CreatedAt: createdDate,
		AgentId:   message.AgentId,
		ChatId:    message.ChatId,
	}

	go chatDb.FetchAIResponse(message.AgentId, chat.Agent.Key, message.Content, message.SystemPrompt, client.Token, message.ChatId, message.Namespaces, textChan)
	log.Println("Fetching...")

	for {
		text := <-textChan

		if text == "____break____" {
			close(textChan)
			go func() {
				log.Println("Saving...")
				chatDb.SaveMessage(&generatedMessage, chat.Agent.Key, message.Content, client.Token)
				chatDb.UpdateChatLastMessage(&generatedMessage)
			}()

			break
		}

		streamedMessage := chatDomain.Message{
			ID:        tempId,
			Content:   text,
			Type:      "bot",
			Sender:    senderId,
			CreatedAt: createdDate,
			AgentId:   message.AgentId,
			ChatId:    message.ChatId,
		}

		data := map[string]any{
			"event": "aiMessage",
			"data":  streamedMessage,
		}

		generatedMessage.Content = generatedMessage.Content + text

		client.sendResponse(data)
	}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
