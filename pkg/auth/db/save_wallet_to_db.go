package auth

import (
	"log"
	"time"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveWalletToDb(email string, wallet *auth.WalletResponse) {
	pool := config.GetPool()
	ctx := getContext()

	now := time.Now()

	_, err := pool.Exec(ctx,
		`INSERT INTO wallets (address, private, email, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5)
		 ON CONFLICT (email) DO UPDATE SET address = $1, private = $2, updated_at = $5`,
		wallet.Address, wallet.KeyPackage[0], email, now, now,
	)

	if err != nil {
		log.Printf("Error trying to save wallet: %v\n", err)
	}
}
