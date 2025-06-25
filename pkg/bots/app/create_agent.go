package bots

import (
	agentsDb "github.com/mmosh-pit/mmosh_backend/pkg/bots/db"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAgent(data agentsDomain.CreateAgentData, userId string) error {

	newId := primitive.NewObjectID()

	data.Id = &newId

	err := agentsDb.SaveAgent(&data)

	if err != nil {
		return err
	}

	ActivateDeactivateAgent(userId, data.Key)

	return nil
}
