package auth

import (
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveLastLogin(id string) {
	pool := config.GetPool()
	ctx := getContext()

	pool.Exec(ctx,
		`UPDATE users SET last_login = $1 WHERE id = $2`,
		time.Now(), id,
	)
}
