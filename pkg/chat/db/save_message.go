package chat

import (
	"log"

	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveMessage(message *chat.Message, chatId *primitive.ObjectID) error {
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

	return err
}
