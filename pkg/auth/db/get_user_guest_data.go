package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetUserGuestData(userId string) *auth.User {
	pool := config.GetPool()
	ctx := getContext()

	row := pool.QueryRow(ctx,
		`SELECT `+selectUserColumns+` FROM users WHERE id = $1`,
		userId,
	)

	result, err := scanUser(row)

	if err != nil {
		return nil
	}

	return &result
}
