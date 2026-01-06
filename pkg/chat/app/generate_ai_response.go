package chat

import (
	"log"
	"math/rand/v2"
	"time"

	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateAIResponse(client *PoolClient, message *chatDomain.Message) {
	log.Println("0")

	textChan := make(chan string)

	id := primitive.NewObjectID()

	chat, _ := chatDb.GetChatById(message.ChatId)

	go chatDb.SaveMessage(message, message.ChatId, chat.Agent.Key, "", client.Token)
	go chatDb.UpdateChatLastMessage(message)

	index := randRange(0, 2)

	senderId := botParticipantUsers[index].ID

	if message.AgentId != nil {
		senderId = message.AgentId
	}

	loadingMessage := chatDomain.Message{
		ID:        &id,
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
		ID:        &id,
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
				chatDb.SaveMessage(&generatedMessage, message.ChatId, chat.Agent.Key, message.Content, client.Token)
				chatDb.UpdateChatLastMessage(&generatedMessage)
			}()

			break
		}

		streamedMessage := chatDomain.Message{
			ID:        &id,
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
