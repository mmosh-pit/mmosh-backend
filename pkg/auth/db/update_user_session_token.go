package auth

import (
	"encoding/json"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateUserSessionToken(sessions []string, id string) error {
	pool := config.GetPool()
	ctx := getContext()

	sessionsJSON, err := json.Marshal(sessions)
	if err != nil {
		return err
	}

	_, err = pool.Exec(ctx,
		`UPDATE users SET sessions = $1 WHERE id = $2`,
		sessionsJSON, id,
	)

	return err
}
