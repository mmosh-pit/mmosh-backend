package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetUserPrivateKey(userId string) (*authDomain.KeyPair, error) {

	user, err := auth.GetUserById(userId)

	data := &authDomain.KeyPair{}

	if err != nil {
		return data, err
	}

	decryptedKey := common.DecryptPrivateKey(user.PrivateKey)

	data.PrivateKey = decryptedKey
	data.PublicKey = user.Address

	return data, nil
}
