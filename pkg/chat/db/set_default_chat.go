package chat

import (
	"context"
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	agents "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetDefaultChat(user *auth.User, bot string) {

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

	log.Printf("Got bot here: %s\n", bot)

	if bot != "KIN" {
		log.Printf("Going to assign new bot: %s\n", bot)
		var newBot agents.Bot

		err := projectCollection.FindOne(*ctx, bson.D{{
			Key:   "symbol",
			Value: bot,
		}}).Decode(&newBot)

		if err != nil {
			log.Printf("Could not setup chat for user, chat not retrieved: %v\n", err)
			return
		}

		newBotId := primitive.NewObjectID()
		newBotChat := chat.Chat{
			ID: &newBotId,
			Agent: &chat.ChatAgent{
				Id:              newBot.Id,
				Name:            newBot.Name,
				Desc:            newBot.Desc,
				Image:           newBot.Image,
				Symbol:          newBot.Symbol,
				Key:             newBot.Key,
				SystemPrompt:    newBot.SystemPrompt,
				CreatorUsername: newBot.CreatorUsername,
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
					ID:      newBot.Id,
					Type:    "bot",
					Picture: newBot.Image,
					Name:    newBot.Name,
				},
			},
		}
		log.Printf("Inserting new bot with: %s\n", bot)
		botCtx := context.Background()
		chatCollection.InsertOne(botCtx, newBotChat)
	}

}
