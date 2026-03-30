package auth

import (
	"encoding/json"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveBlueskyData(data auth.BlueskyUserData, userId string) error {
	pool := config.GetPool()
	ctx := getContext()

	blueskyJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = pool.Exec(ctx,
		`UPDATE users SET bluesky = $1 WHERE id = $2`,
		blueskyJSON, userId,
	)

	return err
}
