package chat

import (
	"fmt"
	"time"

	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
)

var botParticipantUsers []chatDomain.Participant

func SetupBotUsers() {
	botUsers := chatDb.GetBotUsers()

	if len(botUsers) == 0 {
		unclePsyId := fmt.Sprintf("uncle-psy-%d", time.Now().UnixNano())
		auntBeaId := fmt.Sprintf("aunt-bea-%d", time.Now().UnixNano())

		unclePsyParticipant := chatDomain.Participant{
			ID:      unclePsyId,
			Name:    "Uncle psy",
			Type:    "bot",
			Picture: "https://storage.googleapis.com/mmosh-assets/uncle-psy.png",
		}

		auntBeaParticipant := chatDomain.Participant{
			ID:      auntBeaId,
			Name:    "Aunt Bea",
			Type:    "bot",
			Picture: "https://storage.googleapis.com/mmosh-assets/aunt-bea.png",
		}

		chatDb.SaveBotUser(&unclePsyParticipant)
		chatDb.SaveBotUser(&auntBeaParticipant)

		botParticipantUsers = append(botUsers, unclePsyParticipant, auntBeaParticipant)

		return
	}

	botParticipantUsers = botUsers
}
