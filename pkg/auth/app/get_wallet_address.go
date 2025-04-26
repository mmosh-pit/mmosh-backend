package auth

import authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"

func GetWalletAddres(userId string) string {
	user, err := authDb.GetUserById(userId)

	if err != nil {
		return ""
	}

	wallet := authDb.GetWalletByEmail(user.Email)

	return wallet.Address
}
