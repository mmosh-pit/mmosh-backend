package chat

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UnclePsyParticipant = chatDomain.Participant{
	ID:      &primitive.NilObjectID,
	Name:    "Uncle psy",
	Type:    "bot",
	Picture: "https://storage.googleapis.com/mmosh-assets/uncle-psy.png",
}

var AuntBeaParticipant = chatDomain.Participant{
	ID:      &primitive.NilObjectID,
	Name:    "Aunt Bea",
	Type:    "bot",
	Picture: "https://storage.googleapis.com/mmosh-assets/aunt-bea.png",
}

func CreateChat(ownerId *primitive.ObjectID, user *auth.User) *chatDomain.Chat {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	_id := primitive.NewObjectID()

	userParticipantId := primitive.NewObjectID()

	userParticipant := chatDomain.Participant{
		ID:      &userParticipantId,
		Type:    "user",
		Name:    user.Name,
		Picture: "https://storage.googleapis.com/mmosh-assets/avatar_placeholder.png",
	}

	newChat := chatDomain.Chat{
		ID:           &_id,
		Participants: []chatDomain.Participant{userParticipant, UnclePsyParticipant, AuntBeaParticipant},
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
