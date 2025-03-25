package chat

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"time"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateAIResponse(client *PoolClient, message *chatDomain.Message) {
	log.Println("0")
	user, err := auth.GetUserById(message.Sender.Hex())

	if err != nil {
		data := map[string]string{
			"event": "messageError",
			"data":  "error-generating-response",
		}

		encoded, _ := json.Marshal(data)

		client.sendResponse(encoded)
		return
	}

	log.Println("1")

	textChan := make(chan string)

	id := primitive.NewObjectID()

	chatDb.SaveMessage(message, message.ChatId)
	chatDb.UpdateChatLastMessage(message)

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

	data := map[string]interface{}{
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

	go chatDb.FetchAIResponse(user.Name, message.Content, message.SystemPrompt, message.Namespaces, textChan)

	for {
		text := <-textChan

		if text == "____break____" {
			close(textChan)
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

		data := map[string]interface{}{
			"event": "aiMessage",
			"data":  streamedMessage,
		}

		generatedMessage.Content = generatedMessage.Content + text

		client.sendResponse(data)
	}

	chatDb.SaveMessage(&generatedMessage, message.ChatId)
	chatDb.UpdateChatLastMessage(&generatedMessage)
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
