package auth

import (
	"log"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetUserBySubscription(id string) (*authDomain.User, error) {
	pool := config.GetPool()
	ctx := getContext()

	row := pool.QueryRow(ctx,
		`SELECT `+selectUserColumns+` FROM users WHERE subscription->>'subscription_id' = $1`,
		id,
	)

	result, err := scanUser(row)

	if err != nil {
		log.Printf("No user found with subscription %s: %v\n", id, err)
		return nil, common.UserNotExistsErr
	}

	return &result, nil
}
