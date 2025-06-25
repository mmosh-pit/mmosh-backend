package bots

import (
	agentsDb "github.com/mmosh-pit/mmosh_backend/pkg/bots/db"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
)

func GetActiveAgents(userId string) ([]agentsDomain.ActivatedAgentResponse, error) {

	_, err := auth.GetUserById(userId)

	if err != nil {
		return nil, agentsDomain.ErrUserNotFound
	}

	agents := agentsDb.GetActiveAgents(userId)

	return agents, nil
}
