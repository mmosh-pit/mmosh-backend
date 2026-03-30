package chat

import (
	"context"
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetActiveChats(ownerId string) []chatDomain.Chat {
	pool := config.GetPool()
	ctx := context.Background()

	var chats []chatDomain.Chat

	rows, err := pool.Query(ctx,
		`SELECT `+chatSelectColumns+` FROM chats WHERE owner = $1`,
		ownerId,
	)

	if err != nil {
		log.Printf("Got error trying to retrieve active chats: %v\n", err)
		return chats
	}
	defer rows.Close()

	return scanChatRows(rows)
}
