package bots

import (
	"log"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetActiveAgents(userId string) []agentsDomain.ActivatedAgentResponse {

	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-activated-agents")

	resultingAgents := []agentsDomain.ActivatedAgentResponse{}

	res, err := collection.Find(*ctx, bson.D{{Key: "userId", Value: userId}}, &options.FindOptions{
		Projection: map[string]any{
			"agentId": 1,
		},
	})

	if err != nil {
		log.Printf("Could not get activated agents: %v\n", err)
		return resultingAgents
	}

	for res.Next(*ctx) {
		var agent agentsDomain.ActivatedAgentResponse

		if err := res.Decode(&agent); err != nil {
			log.Printf("Error trying to decode activated agent: %v\n", err)

			continue
		}

		resultingAgents = append(resultingAgents, agent)
	}

	return resultingAgents
}
