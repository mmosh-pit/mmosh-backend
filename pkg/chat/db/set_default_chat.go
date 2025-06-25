package chat

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	agents "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetDefaultChat(user *auth.User) {

	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	database := client.Database(databaseName)

	projectCollection := database.Collection("mmosh-app-project")
	chatCollection := database.Collection("chats")

	var defaultBot agents.Bot

	err := projectCollection.FindOne(*ctx, bson.D{{
		Key:   "symbol",
		Value: "KIN",
	}}).Decode(&defaultBot)

	if err != nil {
		log.Printf("Could not setup default chat for user, chat not retrieved: %v\n", err)
		return
	}

	id := primitive.NewObjectID()

	userPicture := ""

	if user.Profile.Image != "" {
		userPicture = user.Profile.Image
	} else if user.GuestData.Picture != "" {
		userPicture = user.GuestData.Picture
	}

	newChat := chat.Chat{
		ID: &id,
		Agent: &chat.ChatAgent{
			Id:              defaultBot.Id,
			Name:            defaultBot.Name,
			Desc:            defaultBot.Desc,
			Image:           defaultBot.Image,
			Symbol:          defaultBot.Symbol,
			Key:             defaultBot.Key,
			SystemPrompt:    defaultBot.SystemPrompt,
			CreatorUsername: defaultBot.CreatorUsername,
		},
		Owner:    user.ID,
		Messages: []chat.Message{},
		Participants: []chat.Participant{
			{
				ID:      user.ID,
				Name:    user.Name,
				Type:    "user",
				Picture: userPicture,
			},
			{
				ID:      defaultBot.Id,
				Type:    "bot",
				Picture: defaultBot.Image,
				Name:    defaultBot.Name,
			},
		},
	}

	chatCollection.InsertOne(*ctx, newChat)

}
