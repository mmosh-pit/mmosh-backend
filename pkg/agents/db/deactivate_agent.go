package agents

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func DeactivateAgent(userId, agentId string) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-activated-agents")

	_, _ = collection.DeleteOne(*ctx, bson.D{{Key: "agentId", Value: agentId}, {Key: "userId", Value: userId}})
}
