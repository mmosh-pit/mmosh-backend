package bots

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	bots "github.com/mmosh-pit/mmosh_backend/pkg/bots/db"
	botsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
)

func GetMyBots(userId string) *[]botsDomain.Bot {
	user, err := auth.GetUserById(userId)

	if err != nil {
		log.Printf("Got error trying to get my bots: %v\n", err)
		return &[]botsDomain.Bot{}
	}

	bots, err := bots.GetMyBots(user.ProfileNFT)

	if err != nil {
		return &[]botsDomain.Bot{}
	}

	return bots
}
