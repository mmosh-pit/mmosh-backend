package chat

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateChat(ownerId *primitive.ObjectID, user *auth.User) *chatDomain.Chat {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	_id := primitive.NewObjectID()

	userParticipant := chatDomain.Participant{
		ID:      ownerId,
		Type:    "user",
		Name:    user.Name,
		Picture: "https://storage.googleapis.com/mmosh-assets/avatar_placeholder.png",
	}

	botParticipants := GetBotUsers()

	participants := botParticipants

	participants = append(participants, userParticipant)

	newChat := chatDomain.Chat{
		ID:           &_id,
		Participants: participants,
		Messages:     []chatDomain.Message{},
		Owner:        ownerId,
	}

	res, err := collection.InsertOne(*ctx, newChat)

	if err != nil {
		log.Printf("Got error creating a new chat: %v\n", err)
		return nil
	}

	id := res.InsertedID.(primitive.ObjectID)

	newChat.ID = &id

	return &newChat
}
