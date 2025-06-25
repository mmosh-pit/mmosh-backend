package bots

import (
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func ActivateAgent(userId, agentId string) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-activated-agents")

	agentData := agentsDomain.ActivatedAgent{
		UserId:  userId,
		AgentId: agentId,
	}

	_, _ = collection.InsertOne(*ctx, agentData)
}
