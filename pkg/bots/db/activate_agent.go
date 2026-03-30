package bots

import (
	"context"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func ActivateAgent(userId, agentId string) {
	pool := config.GetPool()
	ctx := context.Background()

	pool.Exec(ctx,
		`INSERT INTO activated_agents (user_id, agent_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`,
		userId, agentId,
	)
}
