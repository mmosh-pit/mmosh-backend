package auth

import (
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	utils "github.com/mmosh-pit/mmosh_backend/pkg/auth/utils"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
)

type LoginResponse struct {
	Token *string          `json:"token"`
	User  *authDomain.User `json:"user"`
}

func Login(data authDomain.LoginParams) (*LoginResponse, error) {
	user, err := authDb.GetUserByHandle(data.Handle)

	if err != nil {
		return nil, err
	}

	success, verifyPasswordErr := utils.VerifyPassword(data.Password, user.Password)

	if verifyPasswordErr != nil {
		return nil, err
	}

	if !success {
		return nil, common.InvalidPasswordErr
	}

	token, err := utils.GenerateSessionToken(user.Sessions)

	if err != nil {
		return nil, err
	}

	err = authDb.UpdateUserSessionToken(append(user.Sessions, *token), user.ID)

	if err != nil {
		return nil, err
	}

	user.Password = ""
	user.Sessions = []string{}

	CreateWallet(user.Email)

	return &LoginResponse{
		Token: token,
		User:  &user,
	}, nil
}
