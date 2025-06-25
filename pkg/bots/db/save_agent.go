package bots

import (
	"log"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveAgent(data *agentsDomain.CreateAgentData) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-project")

	_, err := collection.InsertOne(*ctx, data)

	if err != nil {
		log.Printf("Could not save agent: %v\n", err)
	}

	return err
}
