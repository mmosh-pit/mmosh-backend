package auth

import (
	"encoding/json"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveTelegramData(data auth.TelegramUserData, userId string) error {
	pool := config.GetPool()
	ctx := getContext()

	telegramJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = pool.Exec(ctx,
		`UPDATE users SET telegram = $1 WHERE id = $2`,
		telegramJSON, userId,
	)

	return err
}
