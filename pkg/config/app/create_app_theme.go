package app

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	configDb "github.com/mmosh-pit/mmosh_backend/pkg/config/db"
	configDomain "github.com/mmosh-pit/mmosh_backend/pkg/config/domain"
)

func CreateAppTheme(data configDomain.AppTheme, userId string) error {
	user, err := auth.GetUserById(userId)

	if err != nil {
		return configDomain.ErrUserNotAuthorized
	}

	if user.Role != "wizard" {
		return configDomain.ErrUserNotAuthorized
	}

	configDb.CreateAppTheme(data)

	return nil
}
