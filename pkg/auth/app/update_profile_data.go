package auth

import (
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
)

func UpdateProfileData(params authDomain.User, userId string) error {
	err := authDb.UpdateProfileData(params, userId)

	authDb.UpdateUserOnboardingStatus(userId, 4)

	return err
}
