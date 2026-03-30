package wallet

import (
	"context"
	"errors"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/jackc/pgx/v5"
)

func GetAssociatedEmailByWallet(walletAddr string) string {
	pool := config.GetPool()
	ctx := context.Background()

	var email string

	err := pool.QueryRow(ctx,
		`SELECT email FROM wallets WHERE address = $1`,
		walletAddr,
	).Scan(&email)

	if errors.Is(err, pgx.ErrNoRows) || err != nil {
		return ""
	}

	return email
}
