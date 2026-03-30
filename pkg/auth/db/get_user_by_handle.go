package auth

import (
	"log"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetUserByHandle(handle string) (authDomain.User, error) {
	pool := config.GetPool()
	ctx := getContext()

	row := pool.QueryRow(ctx,
		`SELECT `+selectUserColumns+` FROM users WHERE email = $1`,
		handle,
	)

	result, err := scanUser(row)

	if err != nil {
		log.Printf("No user found with handle %s: %v\n", handle, err)
		return result, common.UserNotExistsErr
	}

	return result, nil
}
