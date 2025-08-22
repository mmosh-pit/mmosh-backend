package bots

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	bots "github.com/mmosh-pit/mmosh_backend/pkg/bots/db"
	botsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
)

func GetBots(userId, search string, page int) []botsDomain.Bot {

	user, err := auth.GetUserById(userId)

	if err != nil {
		log.Printf("Returning... %s\n", userId)
		return []botsDomain.Bot{}
	}

	username := user.Profile.Username

	return bots.GetBots(search, username, int64(page), user.Role == "wizard")
}
