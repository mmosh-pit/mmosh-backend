package bots

import (
	"context"
	"log"

	botsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetMyBots(creatorUsername string) (*[]botsDomain.Bot, error) {
	pool := config.GetPool()
	ctx := context.Background()

	rows, err := pool.Query(ctx,
		`SELECT `+botSelectColumns+`
		 FROM bots
		 WHERE creator_username = $1
		 ORDER BY created_at DESC
		 LIMIT 100`,
		creatorUsername,
	)

	if err != nil {
		log.Printf("Error in GetMyBots: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var results []botsDomain.Bot

	for rows.Next() {
		var bot botsDomain.Bot
		if err := scanBot(rows, &bot); err != nil {
			log.Printf("Error decoding results: %v\n", err)
			continue
		}
		results = append(results, bot)
	}

	return &results, nil
}
