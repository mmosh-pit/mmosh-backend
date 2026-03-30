package chat

import (
	"context"
	"log"
	"time"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetChatLastMessages(chatId string) []chatDomain.Message {
	pool := config.GetPool()
	ctx := context.Background()

	result := []chatDomain.Message{}

	rows, err := pool.Query(ctx,
		`SELECT id, content, type, created_at, sender, agent_id, chat_id
		 FROM messages
		 WHERE chat_id = $1
		 ORDER BY created_at DESC
		 LIMIT 20`,
		chatId,
	)

	if err != nil {
		log.Printf("[GET CHAT LAST MESSAGES] Got error here: %v\n", err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var msg chatDomain.Message
		var createdAt time.Time

		if err := rows.Scan(&msg.ID, &msg.Content, &msg.Type, &createdAt, &msg.Sender, &msg.AgentId, &msg.ChatId); err != nil {
			log.Printf("[GET CHAT LAST MESSAGES] Error decoding message: %v\n", err)
			continue
		}

		msg.CreatedAt = createdAt
		result = append(result, msg)
	}

	return result
}
