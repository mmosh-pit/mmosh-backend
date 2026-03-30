package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func AddReferrerToUser(username string, userId string) error {
	pool := config.GetPool()
	ctx := getContext()

	_, err := pool.Exec(ctx,
		`UPDATE users SET referred_by = $1 WHERE id = $2`,
		username, userId,
	)

	return err
}
