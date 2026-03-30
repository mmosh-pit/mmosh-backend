package chat

import (
	"context"
	"encoding/json"
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func CreateChat(user *auth.User) *chatDomain.Chat {
	pool := config.GetPool()
	ctx := context.Background()

	userParticipant := chatDomain.Participant{
		ID:      user.ID,
		Type:    "user",
		Name:    user.Name,
		Picture: "https://storage.googleapis.com/mmosh-assets/avatar_placeholder.png",
	}

	botParticipants := GetBotUsers()
	participants := append(botParticipants, userParticipant)

	participantsJSON, _ := json.Marshal(participants)

	var newChatId string
	err := pool.QueryRow(ctx,
		`INSERT INTO chats (owner, participants) VALUES ($1, $2) RETURNING id`,
		user.ID, participantsJSON,
	).Scan(&newChatId)

	if err != nil {
		log.Printf("Got error creating a new chat: %v\n", err)
		return nil
	}

	return &chatDomain.Chat{
		ID:           newChatId,
		Participants: participants,
		Messages:     []chatDomain.Message{},
		Owner:        user.ID,
	}
}
