package bots

import (
	agentsDb "github.com/mmosh-pit/mmosh_backend/pkg/bots/db"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
)

func GetAgents() []agentsDomain.Bot {

	return agentsDb.GetAgents()
}
