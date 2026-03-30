package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func RemoveUserToken(userId string, token string) error {
	pool := config.GetPool()
	ctx := getContext()

	_, err := pool.Exec(ctx,
		`UPDATE users
		 SET sessions = (
		   SELECT COALESCE(jsonb_agg(val), '[]'::jsonb)
		   FROM jsonb_array_elements(sessions) AS val
		   WHERE val <> to_jsonb($1::text)
		 )
		 WHERE id = $2`,
		token, userId,
	)

	return err
}
