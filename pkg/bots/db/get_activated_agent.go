package bots

import (
	"context"
	"errors"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/jackc/pgx/v5"
)

func GetActivatedAgent(userId, agentId string) (*agentsDomain.ActivatedAgent, error) {
	pool := config.GetPool()
	ctx := context.Background()

	var res agentsDomain.ActivatedAgent

	err := pool.QueryRow(ctx,
		`SELECT user_id, agent_id FROM activated_agents WHERE user_id = $1 AND agent_id = $2`,
		userId, agentId,
	).Scan(&res.UserId, &res.AgentId)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return &res, err
}
