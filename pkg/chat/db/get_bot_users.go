package chat

import (
	"context"
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetBotUsers() []chatDomain.Participant {
	pool := config.GetPool()
	ctx := context.Background()

	var resultingUsers []chatDomain.Participant

	rows, err := pool.Query(ctx, `SELECT id, name, type, picture FROM chat_bots`)

	if err != nil {
		return resultingUsers
	}
	defer rows.Close()

	for rows.Next() {
		var user chatDomain.Participant
		if err := rows.Scan(&user.ID, &user.Name, &user.Type, &user.Picture); err != nil {
			log.Printf("Error decoding chat bot participant: %v\n", err)
			continue
		}
		resultingUsers = append(resultingUsers, user)
	}

	return resultingUsers
}
