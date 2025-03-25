package agents

import (
	agentsDb "github.com/mmosh-pit/mmosh_backend/pkg/agents/db"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/agents/domain"
)

func GetAgents() []agentsDomain.Agent {

	return agentsDb.GetAgents()
}
