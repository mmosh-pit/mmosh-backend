package agents

import (
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/agents/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAgentByKey(agentId string) (*agentsDomain.Agent, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-project")

	var agent agentsDomain.Agent

	err := collection.FindOne(*ctx, bson.D{{Key: "key", Value: agentId}}).Decode(&agent)

	return &agent, err
}
