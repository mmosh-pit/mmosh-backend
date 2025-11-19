package chat

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAgentChat(ownerId *primitive.ObjectID, user *auth.User, agent *agentsDomain.Bot) *chatDomain.Chat {
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

	agentParticipant := chatDomain.Participant{
		ID:      agent.Id,
		Name:    agent.Name,
		Type:    "bot",
		Picture: agent.Image,
	}

	participants := []chatDomain.Participant{agentParticipant, userParticipant}

	newChat := chatDomain.Chat{
		ID:           &_id,
		Participants: participants,
		Messages:     []chatDomain.Message{},
		Owner:        ownerId,
		Agent: &chatDomain.ChatAgent{
			Id:              agent.Id,
			Name:            agent.Name,
			Desc:            agent.Desc,
			Image:           agent.Image,
			Symbol:          agent.Symbol,
			Key:             agent.Key,
			SystemPrompt:    agent.SystemPrompt,
			CreatorUsername: agent.CreatorUsername,
			Type:            agent.Type,
			DefaultModel:    agent.DefaultModel,
		},
		Deactivated: false,
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
