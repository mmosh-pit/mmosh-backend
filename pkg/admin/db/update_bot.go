package db

import (
	"context"
	"log"

	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateBot(payload adminDomain.UpdateBotPayload) {
	pool := config.GetPool()
	ctx := context.Background()

	_, err := pool.Exec(ctx,
		`UPDATE bots SET name = $1, symbol = $2, default_model = $3, deactivated = $4 WHERE id = $5`,
		payload.Name, payload.Symbol, payload.DefaultModel, payload.Deactivated, payload.ID,
	)

	if err != nil {
		log.Printf("Error updating BOT: %v\n", err)
	}
}
