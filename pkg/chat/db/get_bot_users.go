package chat

import (
	"log"

	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBotUsers() []chat.Participant {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chat-bots")

	var resultingUsers []chat.Participant

	res, err := collection.Find(*ctx, bson.D{})

	if err != nil {
		return resultingUsers
	}

	for res.Next(*ctx) {
		var user chat.Participant

		if err := res.Decode(&user); err != nil {
			log.Printf("Error decoding chat bot participant: %v\n", err)
			continue
		}

		resultingUsers = append(resultingUsers, user)
	}

	return resultingUsers
}
