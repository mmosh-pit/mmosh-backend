package subscriptions

import (
	"context"
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func DeleteUserSubscription(userId string, productId string) error {
	pool := config.GetPool()
	ctx := context.Background()

	log.Printf("Deleting subscription... %s", userId)

	_, err := pool.Exec(ctx,
		`UPDATE users SET subscription = NULL WHERE id = $1 AND subscription->>'product_id' = $2`,
		userId, productId,
	)

	return err
}
