package chat

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrCreateChatForUser(userId string) *chatDomain.Chat {
	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Printf("Invalid Object ID %s\n", userId)
		return nil
	}

	chat, err := chatDb.GetChat(&objectId)

	if err == common.ChatNotExistsErr {

		user, err := auth.GetUserById(userId)

		if err != nil {
			log.Printf("Could not create chat: %v\n", err)
			return nil
		}

		return chatDb.CreateChat(&objectId, &user)
	}

	return chat
}
