package chat

import (
	"log"

	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveBotUser(user *chat.Participant) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chat-bots")

	_, err := collection.InsertOne(*ctx, user)

	if err != nil {
		log.Printf("Error trying to save bot participants: %v\n", err)
		return err
	}

	return nil
}
