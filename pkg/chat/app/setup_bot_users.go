package chat

import (
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var botParticipantUsers []chatDomain.Participant

func SetupBotUsers() {
	botUsers := chatDb.GetBotUsers()

	if len(botUsers) == 0 {
		unclePsyId := primitive.NewObjectID()

		auntBeaId := primitive.NewObjectID()

		var unclePsyParticipant = chatDomain.Participant{
			ID:      &unclePsyId,
			Name:    "Uncle psy",
			Type:    "bot",
			Picture: "https://storage.googleapis.com/mmosh-assets/uncle-psy.png",
		}

		var auntBeaParticipant = chatDomain.Participant{
			ID:      &auntBeaId,
			Name:    "Aunt Bea",
			Type:    "bot",
			Picture: "https://storage.googleapis.com/mmosh-assets/aunt-bea.png",
		}

		chatDb.SaveBotUser(&unclePsyParticipant)
		chatDb.SaveBotUser(&auntBeaParticipant)

		botParticipantUsers = append(botUsers, unclePsyParticipant)
		botParticipantUsers = append(botUsers, auntBeaParticipant)

		return
	}

	botParticipantUsers = botUsers
}
