package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func DeleteTelegramData(userId string) error {
	pool := config.GetPool()
	ctx := getContext()

	_, err := pool.Exec(ctx,
		`UPDATE users SET telegram = NULL WHERE id = $1`,
		userId,
	)

	return err
}
