package bots

import (
	"context"
	"log"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetActiveAgents(userId string) []agentsDomain.ActivatedAgentResponse {
	pool := config.GetPool()
	ctx := context.Background()

	resultingAgents := []agentsDomain.ActivatedAgentResponse{}

	rows, err := pool.Query(ctx,
		`SELECT agent_id FROM activated_agents WHERE user_id = $1`,
		userId,
	)

	if err != nil {
		log.Printf("Could not get activated agents: %v\n", err)
		return resultingAgents
	}
	defer rows.Close()

	for rows.Next() {
		var agent agentsDomain.ActivatedAgentResponse
		if err := rows.Scan(&agent.AgentId); err != nil {
			log.Printf("Error trying to decode activated agent: %v\n", err)
			continue
		}
		resultingAgents = append(resultingAgents, agent)
	}

	return resultingAgents
}
