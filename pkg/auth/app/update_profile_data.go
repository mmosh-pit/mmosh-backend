package auth

import (
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProfileData(params authDomain.User, userId string) error {

	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	err = authDb.UpdateProfileData(params, userIdBson)

	authDb.UpdateUserOnboardingStatus(&userIdBson, 4)

	return err
}
