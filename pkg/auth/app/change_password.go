package auth

import (
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	utils "github.com/mmosh-pit/mmosh_backend/pkg/auth/utils"
	commonDomain "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
)

type ChangePasswordParams struct {
	Code     int    `json:"code"`
	Password string `json:"password"`
}

func ChangePassword(params ChangePasswordParams) error {

	existingCode := authDb.GetTemporalCode(params.Code)

	if existingCode == nil {
		return commonDomain.InvalidCodeErr
	}

	existingUser, err := authDb.GetUserByEmail(existingCode.Email)

	if err != nil || existingUser.ID == nil {
		return authDomain.ErrUserNotExists
	}

	password, err := utils.EncryptPassword(params.Password)

	return authDb.UpdateUserPassword(existingUser.ID, password)
}
