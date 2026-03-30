package chat

import (
	"context"
	"encoding/json"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateChatLastMessage(message *chatDomain.Message) {
	pool := config.GetPool()
	ctx := context.Background()

	msgJSON, _ := json.Marshal(message)

	pool.Exec(ctx,
		`UPDATE chats SET last_message = $1 WHERE id = $2`,
		msgJSON, message.ChatId,
	)
}
