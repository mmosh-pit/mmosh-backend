package chat

import (
	"context"
	"encoding/json"
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func CreateAgentChat(ownerId string, user *auth.User, agent *agentsDomain.Bot) *chatDomain.Chat {
	pool := config.GetPool()
	ctx := context.Background()

	chatAgent := &chatDomain.ChatAgent{
		Id:              agent.Id,
		Name:            agent.Name,
		Desc:            agent.Desc,
		Image:           agent.Image,
		Symbol:          agent.Symbol,
		Key:             agent.Key,
		SystemPrompt:    agent.SystemPrompt,
		CreatorUsername: agent.CreatorUsername,
		Type:            agent.Type,
		DefaultModel:    agent.DefaultModel,
	}

	participants := []chatDomain.Participant{
		{
			ID:      agent.Id,
			Name:    agent.Name,
			Type:    "bot",
			Picture: agent.Image,
		},
		{
			ID:      ownerId,
			Type:    "user",
			Name:    user.Name,
			Picture: "https://storage.googleapis.com/mmosh-assets/avatar_placeholder.png",
		},
	}

	agentJSON, _ := json.Marshal(chatAgent)
	participantsJSON, _ := json.Marshal(participants)

	var newChatId string
	err := pool.QueryRow(ctx,
		`INSERT INTO chats (owner, chat_agent, participants) VALUES ($1, $2, $3) RETURNING id`,
		ownerId, agentJSON, participantsJSON,
	).Scan(&newChatId)

	if err != nil {
		log.Printf("Got error creating a new chat: %v\n", err)
		return nil
	}

	return &chatDomain.Chat{
		ID:           newChatId,
		Participants: participants,
		Messages:     []chatDomain.Message{},
		Owner:        ownerId,
		Agent:        chatAgent,
		Deactivated:  false,
	}
}
