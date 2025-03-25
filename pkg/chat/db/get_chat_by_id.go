package chat

import (
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetChatById(id *primitive.ObjectID) (*chatDomain.Chat, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	var result chatDomain.Chat

	err := collection.FindOne(*ctx, bson.D{{Key: "_id", Value: id}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return nil, common.ChatNotExistsErr
	}

	if err != nil {
		log.Printf("Got error: %v\n", err)
		return nil, err
	}

	return &result, nil
}
