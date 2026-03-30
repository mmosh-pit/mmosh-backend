package bots

import (
	"context"
	"errors"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/jackc/pgx/v5"
)

func GetAgentByKey(agentKey string) (*agentsDomain.Bot, error) {
	pool := config.GetPool()
	ctx := context.Background()

	var agent agentsDomain.Bot

	err := pool.QueryRow(ctx,
		`SELECT id, name, symbol, description, key, image, creator_username, type, system_prompt, default_model, deactivated, created_at
		 FROM bots WHERE key = $1`,
		agentKey,
	).Scan(
		&agent.Id, &agent.Name, &agent.Symbol, &agent.Desc, &agent.Key, &agent.Image,
		&agent.CreatorUsername, &agent.Type, &agent.SystemPrompt, &agent.DefaultModel,
		&agent.Deactivated, &agent.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return &agent, err
}
