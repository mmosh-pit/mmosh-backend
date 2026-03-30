package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	bots "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetAllBots(page int64, search string) *[]bots.Bot {
	pool := config.GetPool()
	ctx := context.Background()

	result := []bots.Bot{}

	var (
		rows pgx.Rows
		err  error
	)

	const selectCols = `id, name, symbol, description, key, image, creator_username, type, system_prompt, default_model, deactivated, created_at`

	if search != "" {
		pattern := "%" + search + "%"
		rows, err = pool.Query(ctx,
			`SELECT `+selectCols+`
			 FROM bots
			 WHERE name ILIKE $1 OR symbol ILIKE $1 OR description ILIKE $1
			 ORDER BY created_at DESC
			 LIMIT 20 OFFSET $2`,
			pattern, page*20,
		)
	} else {
		rows, err = pool.Query(ctx,
			`SELECT `+selectCols+`
			 FROM bots
			 ORDER BY created_at DESC
			 LIMIT 20 OFFSET $1`,
			page*20,
		)
	}

	if err != nil {
		log.Printf("[ADMIN/GET BOTS] Got error here: %v\n", err)
		return &result
	}
	defer rows.Close()

	for rows.Next() {
		var bot bots.Bot
		if err := rows.Scan(
			&bot.Id, &bot.Name, &bot.Symbol, &bot.Desc, &bot.Key, &bot.Image,
			&bot.CreatorUsername, &bot.Type, &bot.SystemPrompt, &bot.DefaultModel,
			&bot.Deactivated, &bot.CreatedAt,
		); err != nil {
			log.Printf("[ADMIN/GET BOTS] Error scanning bot: %v\n", err)
			continue
		}
		result = append(result, bot)
	}

	return &result
}
