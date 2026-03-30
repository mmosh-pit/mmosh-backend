package auth

import (
	"log"

	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/utils"
)

func AddBlueskyData(data authDomain.BlueskyUserData, userId string) error {
	isValid := auth.IsValidConnection(data.Handle, data.Password)

	if !isValid {
		return authDomain.ErrInvalidBluesky
	}

	err := authDb.SaveBlueskyData(data, userId)

	if err != nil {
		log.Printf("Got error saving bluesky data: %v\n", err)
		return err
	}

	return nil
}
