package auth

import (
	"log"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetUserById(id string) (authDomain.User, error) {
	pool := config.GetPool()
	ctx := getContext()

	row := pool.QueryRow(ctx,
		`SELECT `+selectUserColumns+` FROM users WHERE id = $1`,
		id,
	)

	result, err := scanUser(row)

	if err != nil {
		log.Printf("No user found with id %s: %v\n", id, err)
		return result, err
	}

	result.Password = ""
	result.Sessions = nil

	return result, nil
}
