package bots

import (
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAgentByKey(agentId string) (*agentsDomain.Bot, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-project")

	var agent agentsDomain.Bot

	err := collection.FindOne(*ctx, bson.D{{Key: "key", Value: agentId}}).Decode(&agent)

	return &agent, err
}
