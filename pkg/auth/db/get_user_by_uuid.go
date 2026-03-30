package auth

import (
	"log"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetUserByUuidId(id string) (authDomain.User, error) {
	pool := config.GetPool()
	ctx := getContext()

	row := pool.QueryRow(ctx,
		`SELECT `+selectUserColumns+` FROM users WHERE uuid = $1`,
		id,
	)

	result, err := scanUser(row)

	if err != nil {
		log.Printf("No user found with uuid %s: %v\n", id, err)
		return result, err
	}

	return result, nil
}
