package bots

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	agentsDb "github.com/mmosh-pit/mmosh_backend/pkg/bots/db"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
)

func GetActiveAgents(userId, search string) ([]agentsDomain.ActivatedAgentResponse, error) {

	_, err := auth.GetUserById(userId)

	if err != nil {
		return nil, agentsDomain.ErrUserNotFound
	}

	agents := agentsDb.GetActiveAgents(userId)

	return agents, nil
}
