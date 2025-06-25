package bots

import (
	agentsDb "github.com/mmosh-pit/mmosh_backend/pkg/bots/db"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
)

func ActivateDeactivateAgent(userId, agentId string) error {
	user, err := authDb.GetUserById(userId)

	if err != nil {
		return agentsDomain.ErrUserNotFound
	}

	// expiresAt := time.Unix(0, user.Subscription.ExpiresAt)
	//
	// if user.Subscription.ProductId == "" || expiresAt.Before(time.Now()) {
	// 	return agentsDomain.ErrUserNotSubscribed
	// }

	agent, err := agentsDb.GetAgentByKey(agentId)

	if err != nil {
		return agentsDomain.ErrAgentNotExists
	}

	activatedAgent, err := agentsDb.GetActivatedAgent(userId, agentId)

	if err != nil {
		return agentsDomain.ErrAgentNotExists
	}

	chat := chatDb.GetChatByAgentAndUser(userId, agentId)

	if activatedAgent == nil {
		agentsDb.ActivateAgent(userId, agentId)

		if chat == nil {
			chatDb.CreateAgentChat(user.ID, &user, agent)
		} else {
			chatDb.ActivateChat(userId, agentId)
		}

	} else {
		if chat != nil {
			chatDb.DeactivateChat(userId, agentId)
		}

		agentsDb.DeactivateAgent(userId, agentId)
	}

	return nil
}
