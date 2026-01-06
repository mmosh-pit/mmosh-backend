package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveMessage(message *chat.Message, chatId *primitive.ObjectID, chatAgentKey string, userContent string, authToken string) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	filter := bson.D{{Key: "_id", Value: chatId}}

	update := bson.D{{Key: "$push", Value: bson.D{{Key: "messages", Value: message}}}}

	// update := bson.M{"$push": bson.M{"messages": message}}

	_, err := collection.UpdateOne(*ctx, filter, update)

	if err != nil {
		log.Printf("Error trying to save message: %v\n", err)
	}

	data := map[string]any{
		"chatId":       chatId.Hex(),
		"namespaces":   []string{chatAgentKey, "PUBLIC"},
		"systemPrompt": message.SystemPrompt,
		// "chatHistory":  formatChatHistory(messages),
		"agentID":     message.AgentId.Hex(),
		"aiModel":     "gpt-5.1",
		"botContent":  message.Content,
		"userContent": userContent,
	}

	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error trying to encode following payload: %v\n, %v\n", data, err)
		return err
	}

	baseUrl := config.GetAIBaseUrl()

	url := fmt.Sprintf("%s/save-chat", baseUrl)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(encoded))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	httpClient := http.Client{
		Timeout: time.Second * 500,
	}
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Printf("[SAVE MESSAGE] Error in request: %v\n", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)

	log.Printf("[SAVE MESSAGE] response: %v\n", string(body))

	return nil
}
