package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	commonApp "github.com/mmosh-pit/mmosh_backend/pkg/common/app"
)

func AddEarlyAccess(params authDomain.AddEarlyAccessParams) error {
	err := auth.GetEarlyAccessData(params.Email)

	if err != nil {
		return authDomain.ErrEarlyAlreadyRegistered
	}

	user, err := auth.GetUserByEmail(params.Email)

	if err != nil || user.Email == params.Email {
		return authDomain.ErrEarlyAlreadyRegistered
	}

	err = auth.SaveEarlyAccess(params)

	if err != nil {
		return err
	}

	go commonApp.SendKartraNotification("kinship_bots_waitlist", params.Name, "", params.Email)

	return nil
}
