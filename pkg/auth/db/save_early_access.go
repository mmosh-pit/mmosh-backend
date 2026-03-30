package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveEarlyAccess(params auth.AddEarlyAccessParams) error {
	pool := config.GetPool()
	ctx := getContext()

	_, err := pool.Exec(ctx,
		`INSERT INTO early_access (name, email) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING`,
		params.Name, params.Email,
	)

	return err
}
