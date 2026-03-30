package auth

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
)

func AddReferrerUser(data authDomain.AddReferrerParams, userId string) error {
	_, err := auth.GetReferredUser(data.User)

	if err != nil {
		return err
	}

	err = auth.AddReferrerToUser(data.User, userId)

	if err != nil {
		log.Printf("Got error trying to add refer: %v\n", err)
		return err
	}

	auth.UpdateUserOnboardingStatus(userId, 2)

	return nil
}
