package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveTemporalCode(email string, code int) {
	pool := config.GetPool()
	ctx := getContext()

	pool.Exec(ctx,
		`INSERT INTO email_verification (email, code) VALUES ($1, $2)`,
		email, code,
	)
}
