package auth

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetWalletByEmail(email string) *auth.Wallet {
	pool := config.GetPool()
	ctx := getContext()

	var data auth.Wallet

	err := pool.QueryRow(ctx,
		`SELECT address, private, email, created_at, updated_at FROM wallets WHERE email = $1`,
		email,
	).Scan(&data.Address, &data.Private, &data.Email, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		log.Printf("Error querying wallet: %v\n", err)
		return nil
	}

	return &data
}
