package chat

import (
	"errors"
	"log"
	"time"

	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	bots "github.com/mmosh-pit/mmosh_backend/pkg/bots/db"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	wallet "github.com/mmosh-pit/mmosh_backend/pkg/wallet/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrWalletNotFound = errors.New("recipient-wallet-not-found")

type SendBotMessageData struct {
	Sender          string `json:"sender"`
	RecipientWallet string `json:"recipient_wallet"`
	RecipientName   string `json:"recipient_name"`
	Message         string `json:"message"`
	OriginalMessage string `json:"original_message"`
	SentMessageId   string `json:"sent_message_id"`
	BotId           string `json:"bot_id"`
}

func SendBotMessage(params SendBotMessageData) error {
	userEmail := wallet.GetAssociatedEmailByWallet(params.RecipientWallet)

	if userEmail == "" {
		return ErrWalletNotFound
	}

	recipient, err := authDb.GetUserByEmail(userEmail)

	if err != nil {
		return authDomain.ErrUserNotExists
	}

	bot, err := bots.GetAgentByKey(params.BotId)

	if err != nil {
		log.Printf("[SEND BOT MESSAGE] Could not fetch agent by key: %v\n", err)
		return err
	}

	id := primitive.NewObjectID()

	senderId, err := primitive.ObjectIDFromHex(params.Sender)

	if err != nil {
		log.Printf("[SEND BOT MESSAGE] invalid sender ID: %s, %v\n", params.Sender, err)
		return err
	}

	messageData := chatDomain.Message{
		ID:        &id,
		Content:   params.Message,
		Sender:    &senderId,
		CreatedAt: time.Now(),
		Type:      "user",
		AgentId:   bot.Id,
		ChatId:    bot.Id,
	}

	chatDb.SaveMessage(&messageData, messageData.ChatId)

	client := WsPool.Clients[recipient.ID.Hex()]

	if client != nil {
		client.sendResponse(messageData)
	}

	return nil
}
