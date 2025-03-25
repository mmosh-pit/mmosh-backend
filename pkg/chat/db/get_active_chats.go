package chat

import (
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetActiveChats(ownerId string) []chatDomain.Chat {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	var chats []chatDomain.Chat

	parsedOwnerId, err := primitive.ObjectIDFromHex(ownerId)

	if err != nil {
		log.Printf("Invalid owner id: %v, %v\n", ownerId, err)
		return chats
	}

	res, err := collection.Find(*ctx, bson.D{{Key: "owner", Value: parsedOwnerId}})

	if err != nil {
		log.Printf("Got error trying to retrieve active chats: %v\n", err)
		return chats
	}

	for res.Next(*ctx) {
		var chat chatDomain.Chat

		if err := res.Decode(&chat); err != nil {
			log.Printf("Got error decoding a chat: %v\n", err)
			continue
		}

		chats = append(chats, chat)
	}

	return chats
}
