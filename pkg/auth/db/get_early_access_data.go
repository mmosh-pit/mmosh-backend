package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetEarlyAccessData(email string) error {
	pool := config.GetPool()
	ctx := getContext()

	var existingEmail string

	err := pool.QueryRow(ctx,
		`SELECT email FROM early_access WHERE email = $1`,
		email,
	).Scan(&existingEmail)

	if err != nil {
		return nil
	}

	return nil
}
