package chat

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func ActivateChat(userId, agentId string) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	filter := bson.D{{Key: "owner", Value: userId}, {Key: "agent.id", Value: agentId}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "deactivated", Value: false}}}}

	_, err := collection.UpdateOne(*ctx, filter, update)

	return err
}
