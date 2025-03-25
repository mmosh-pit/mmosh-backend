package agents

import (
	"errors"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/agents/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetActivatedAgent(userId, agentId string) (*agentsDomain.ActivatedAgent, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-activated-agents")

	var res agentsDomain.ActivatedAgent

	err := collection.FindOne(*ctx, bson.D{{Key: "userId", Value: userId}, {Key: "agentId", Value: agentId}}).Decode(&res)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	return &res, err
}
