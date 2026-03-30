package chat

import (
	"context"
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/jackc/pgx/v5"
)

func GetChatByAgentAndUser(userId, agentId string) *chatDomain.Chat {
	pool := config.GetPool()
	ctx := context.Background()

	row := pool.QueryRow(ctx,
		`SELECT `+chatSelectColumns+` FROM chats WHERE owner = $1 AND chat_agent->>'id' = $2`,
		userId, agentId,
	)

	result, err := scanChat(row)

	if err == pgx.ErrNoRows {
		return nil
	}

	if err != nil {
		log.Printf("Got error while fetching chat by agent and user: %v\n", err)
		return nil
	}

	return &result
}
