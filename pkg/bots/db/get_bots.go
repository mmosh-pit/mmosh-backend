package bots

import (
	"context"
	"log"

	botsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/jackc/pgx/v5"
)

const botSelectColumns = `id, name, symbol, description, key, image, creator_username, type, system_prompt, default_model, deactivated, created_at`

func scanBot(rows pgx.Rows, bot *botsDomain.Bot) error {
	return rows.Scan(
		&bot.Id, &bot.Name, &bot.Symbol, &bot.Desc, &bot.Key, &bot.Image,
		&bot.CreatorUsername, &bot.Type, &bot.SystemPrompt, &bot.DefaultModel,
		&bot.Deactivated, &bot.CreatedAt,
	)
}

func GetBots(search, username string, page int64, isWizard bool) []botsDomain.Bot {
	pool := config.GetPool()
	ctx := context.Background()

	result := []botsDomain.Bot{}

	var (
		rows pgx.Rows
		err  error
	)

	if search != "" {
		pattern := "%" + search + "%"
		rows, err = pool.Query(ctx,
			`SELECT `+botSelectColumns+`
			 FROM bots
			 WHERE (name ILIKE $1 OR symbol ILIKE $1 OR description ILIKE $1)
			   AND (privacy = 'public' OR privacy IS NULL OR creator_username = $2 OR (privacy = 'secret' AND key = $3))
			 ORDER BY created_at DESC
			 LIMIT 20 OFFSET $4`,
			pattern, username, search, page*20,
		)
	} else if isWizard {
		rows, err = pool.Query(ctx,
			`SELECT `+botSelectColumns+`
			 FROM bots
			 ORDER BY created_at DESC
			 LIMIT 20 OFFSET $1`,
			page*20,
		)
	} else {
		rows, err = pool.Query(ctx,
			`SELECT `+botSelectColumns+`
			 FROM bots
			 WHERE privacy = 'public' OR privacy IS NULL OR creator_username = $1
			 ORDER BY created_at DESC
			 LIMIT 20 OFFSET $2`,
			username, page*20,
		)
	}

	if err != nil {
		log.Printf("[GET BOTS] Got error here: %v\n", err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var bot botsDomain.Bot
		if err := scanBot(rows, &bot); err != nil {
			log.Printf("[GET BOTS] Error decoding bot: %v\n", err)
			continue
		}
		result = append(result, bot)
	}

	return result
}
