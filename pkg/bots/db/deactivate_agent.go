package bots

import (
	"context"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func DeactivateAgent(userId, agentId string) {
	pool := config.GetPool()
	ctx := context.Background()

	pool.Exec(ctx,
		`DELETE FROM activated_agents WHERE user_id = $1 AND agent_id = $2`,
		userId, agentId,
	)
}
