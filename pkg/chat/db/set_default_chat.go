package chat

import (
	"context"
	"encoding/json"
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SetDefaultChat(user *auth.User, bot string) {
	pool := config.GetPool()
	ctx := context.Background()

	insertChat := func(botId, botName, botDesc, botImage, botSymbol, botKey, botSystemPrompt, botCreatorUsername string) {
		chatAgent := chatDomain.ChatAgent{
			Id:              botId,
			Name:            botName,
			Desc:            botDesc,
			Image:           botImage,
			Symbol:          botSymbol,
			Key:             botKey,
			SystemPrompt:    botSystemPrompt,
			CreatorUsername: botCreatorUsername,
		}
		participants := []chatDomain.Participant{
			{
				ID:      user.ID,
				Name:    user.Name,
				Type:    "user",
				Picture: user.Picture,
			},
			{
				ID:      botId,
				Type:    "bot",
				Picture: botImage,
				Name:    botName,
			},
		}

		agentJSON, _ := json.Marshal(chatAgent)
		participantsJSON, _ := json.Marshal(participants)

		_, err := pool.Exec(ctx,
			`INSERT INTO chats (owner, chat_agent, participants) VALUES ($1, $2, $3)`,
			user.ID, agentJSON, participantsJSON,
		)
		if err != nil {
			log.Printf("Could not insert chat: %v\n", err)
		}
	}

	var (
		botId, botName, botDesc, botImage, botSymbol, botKey, botSystemPrompt, botCreatorUsername string
	)

	err := pool.QueryRow(ctx,
		`SELECT id, name, description, image, symbol, key, system_prompt, creator_username FROM bots WHERE symbol = 'CATFAWN'`,
	).Scan(&botId, &botName, &botDesc, &botImage, &botSymbol, &botKey, &botSystemPrompt, &botCreatorUsername)

	if err != nil {
		log.Printf("Could not setup default chat for user, chat not retrieved: %v\n", err)
		return
	}

	insertChat(botId, botName, botDesc, botImage, botSymbol, botKey, botSystemPrompt, botCreatorUsername)

	log.Printf("Got bot here: %s\n", bot)

	if bot != "KIN" {
		log.Printf("Going to assign new bot: %s\n", bot)

		err := pool.QueryRow(ctx,
			`SELECT id, name, description, image, symbol, key, system_prompt, creator_username FROM bots WHERE symbol = $1`,
			bot,
		).Scan(&botId, &botName, &botDesc, &botImage, &botSymbol, &botKey, &botSystemPrompt, &botCreatorUsername)

		if err != nil {
			log.Printf("Could not setup chat for user, chat not retrieved: %v\n", err)
			return
		}

		log.Printf("Inserting new bot with: %s\n", bot)
		insertChat(botId, botName, botDesc, botImage, botSymbol, botKey, botSystemPrompt, botCreatorUsername)
	}
}
