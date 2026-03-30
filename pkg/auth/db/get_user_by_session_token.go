package auth

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetUserBySessionToken(token string) (*authDomain.User, error) {
	pool := config.GetPool()
	ctx := getContext()

	row := pool.QueryRow(ctx,
		`SELECT `+selectUserColumns+` FROM users WHERE sessions @> jsonb_build_array($1::text)`,
		token,
	)

	user, err := scanUser(row)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
