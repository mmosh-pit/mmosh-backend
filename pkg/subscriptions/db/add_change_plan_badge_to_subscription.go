package subscriptions

import (
	"context"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func AddChangePlanBadgeToSubscription(userId string) error {
	pool := config.GetPool()
	ctx := context.Background()

	_, err := pool.Exec(ctx,
		`UPDATE users SET subscription = subscription || '{"changed_plan": true}'::jsonb WHERE id = $1`,
		userId,
	)

	return err
}
