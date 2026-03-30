package bots

import (
	"context"
	"log"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetAgents() []agentsDomain.Bot {
	pool := config.GetPool()
	ctx := context.Background()

	var resultingAgents []agentsDomain.Bot

	rows, err := pool.Query(ctx,
		`SELECT id, name, symbol, description, key, image, creator_username, type, system_prompt, default_model, deactivated, created_at
		 FROM bots`,
	)

	if err != nil {
		log.Printf("Got error returning agents: %v\n", err)
		return resultingAgents
	}
	defer rows.Close()

	for rows.Next() {
		var agent agentsDomain.Bot
		if err := rows.Scan(
			&agent.Id, &agent.Name, &agent.Symbol, &agent.Desc, &agent.Key, &agent.Image,
			&agent.CreatorUsername, &agent.Type, &agent.SystemPrompt, &agent.DefaultModel,
			&agent.Deactivated, &agent.CreatedAt,
		); err != nil {
			log.Printf("Error decoding agent: %v\n", err)
			continue
		}
		resultingAgents = append(resultingAgents, agent)
	}

	return resultingAgents
}
