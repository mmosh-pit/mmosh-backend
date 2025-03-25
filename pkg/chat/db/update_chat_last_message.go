package chat

import (
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateChatLastMessage(message *chat.Message) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	filter := bson.D{{Key: "_id", Value: message.ChatId}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "lastMessage", Value: message}}}}

	collection.UpdateOne(*ctx, filter, update)
}
