package auth

import (
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGuestUserData(data authDomain.GuestUserData, userId string) error {

	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return authDomain.ErrSomethingWentWrong
	}

	err = authDb.AddUserGuestData(data, &userIdBson)

	if err != nil {
		return authDomain.ErrSomethingWentWrong
	}

	authDb.UpdateUserOnboardingStatus(&userIdBson, 4)

	return nil
}
