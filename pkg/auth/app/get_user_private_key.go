package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
)

func GetUserPrivateKey(userId string) (*authDomain.KeyPair, error) {

	_, err := auth.GetUserById(userId)

	data := &authDomain.KeyPair{}

	if err != nil {
		return data, err
	}

	// decryptedKey := common.DecryptPrivateKey(user.PrivateKey)
	//
	// data.PrivateKey = decryptedKey
	// data.PublicKey = user.Address

	return data, nil
}
