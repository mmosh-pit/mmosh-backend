package chat

import (
	"context"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func DeactivateChat(userId, agentId string) error {
	pool := config.GetPool()
	ctx := context.Background()

	_, err := pool.Exec(ctx,
		`UPDATE chats SET deactivated = true WHERE owner = $1 AND chat_agent->>'id' = $2`,
		userId, agentId,
	)

	return err
}
