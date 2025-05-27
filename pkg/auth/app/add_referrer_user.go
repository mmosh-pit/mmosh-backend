package auth

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddReferrerUser(data authDomain.AddReferrerParams, userId string) error {

	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return authDomain.ErrSomethingWentWrong
	}

	_, err = auth.GetReferredUser(data.User)

	if err != nil {
		return err
	}

	err = auth.AddReferrerToUser(data.User, &userIdBson)

	if err != nil {
		log.Printf("Got error trying to add refer: %v\n", err)
		return err
	}

	auth.UpdateUserOnboardingStatus(&userIdBson, 2)

	return nil
}
