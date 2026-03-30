package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateUserPassword(ID string, password string) error {
	pool := config.GetPool()
	ctx := getContext()

	_, err := pool.Exec(ctx,
		`UPDATE users SET password = $1 WHERE id = $2`,
		password, ID,
	)

	return err
}
