package bots

import (
	"log"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAgents() []agentsDomain.Bot {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-project")

	var resultingAgents []agentsDomain.Bot

	res, err := collection.Find(*ctx, bson.D{{}})

	if err != nil {
		log.Printf("Got error returning agents: %v\n", err)
		return resultingAgents
	}

	for res.Next(*ctx) {
		var agent agentsDomain.Bot

		if err := res.Decode(&agent); err != nil {
			log.Printf("Error decoding agent: %v\n", err)
			continue
		}

		resultingAgents = append(resultingAgents, agent)
	}

	return resultingAgents
}
