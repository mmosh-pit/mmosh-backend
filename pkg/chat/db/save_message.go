package chat

import (
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveMessage(message *chat.Message) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	filter := bson.D{{Key: "_id", Value: message.Sender}}

	update := bson.D{{Key: "$push", Value: bson.D{{Key: "messages", Value: message}}}}

	_, err := collection.UpdateOne(*ctx, filter, update)

	return err
}
