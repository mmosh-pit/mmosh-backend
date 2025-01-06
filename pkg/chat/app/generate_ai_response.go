package chat

import (
	"encoding/json"
	"math/rand/v2"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateAIResponse(client *PoolClient, message *chatDomain.Message) {
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

	textChan := make(chan string)

	id := primitive.NewObjectID()

	chatDb.SaveMessage(message, message.Sender)

	index := randRange(0, 2)

	loadingMessage := chatDomain.Message{
		ID:        &id,
		Content:   "",
		Type:      "bot",
		Sender:    botParticipantUsers[index].ID,
		IsLoading: true,
	}

	client.sendResponse(loadingMessage)

	generatedMessage := chatDomain.Message{
		ID:      &id,
		Content: "",
		Type:    "bot",
		Sender:  botParticipantUsers[index].ID,
	}

	go chatDb.FetchAIResponse(user.Name, message.Content, []string{"PUBLIC"}, textChan)

	for {
		text := <-textChan

		if text == "____break____" {
			close(textChan)
			break
		}

		streamedMessage := chatDomain.Message{
			ID:      &id,
			Content: text,
			Type:    "bot",
			Sender:  botParticipantUsers[index].ID,
		}

		data := map[string]interface{}{
			"event": "aiMessage",
			"data":  streamedMessage,
		}

		generatedMessage.Content = generatedMessage.Content + text

		client.sendResponse(data)
	}

	chatDb.SaveMessage(&generatedMessage, message.Sender)
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}
