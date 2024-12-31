package chat

import (
	"encoding/json"

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

	chatDb.SaveMessage(message)

	go chatDb.FetchAIResponse(user.Name, message.Content, []string{"PUBLIC"}, textChan)

	generatedMessage := chatDomain.Message{
		ID:      &id,
		Content: "",
		Type:    "bot",
		Sender:  &primitive.NilObjectID,
	}

	for {
		text := <-textChan

		if text == "____break____" {
			close(textChan)
			break
		}

		data := map[string]interface{}{
			"event": "aiMessage",
			"data": map[string]interface{}{
				"content": text,
				"id":      id,
			},
		}

		generatedMessage.Content = generatedMessage.Content + text

		client.sendResponse(data)
	}

	chatDb.SaveMessage(&generatedMessage)
}
