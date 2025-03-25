package chat

import (
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetChatByAgentAndUser(userId, agentId string) *chatDomain.Chat {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	var result chatDomain.Chat

	err := collection.FindOne(*ctx, bson.D{{Key: "owner", Value: userId}, {Key: "agent.id", Value: agentId}}).Decode(&result)

	if err != nil {
		log.Printf("Got error while fetching chat by agent and user: %v\n", err)

		return nil
	}

	return &result
}
