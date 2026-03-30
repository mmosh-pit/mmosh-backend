package chat

import (
	"context"
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveBotUser(user *chatDomain.Participant) error {
	pool := config.GetPool()
	ctx := context.Background()

	_, err := pool.Exec(ctx,
		`INSERT INTO chat_bots (id, name, type, picture) VALUES ($1, $2, $3, $4)
		 ON CONFLICT (id) DO UPDATE SET name = $2, type = $3, picture = $4`,
		user.ID, user.Name, user.Type, user.Picture,
	)

	if err != nil {
		log.Printf("Error trying to save bot participants: %v\n", err)
		return err
	}

	return nil
}
