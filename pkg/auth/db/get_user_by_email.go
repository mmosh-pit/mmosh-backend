package auth

import (
	"log"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetUserByEmail(email string) (authDomain.User, error) {
	pool := config.GetPool()
	ctx := getContext()

	row := pool.QueryRow(ctx,
		`SELECT `+selectUserColumns+` FROM users WHERE email = $1`,
		email,
	)

	result, err := scanUser(row)

	if err != nil {
		log.Printf("No user found with email %s: %v\n", email, err)
		return result, err
	}

	result.Password = ""

	return result, nil
}
