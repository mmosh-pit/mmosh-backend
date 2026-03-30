package subscriptions

import (
	"context"
	"encoding/json"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateSubscriptionStatus(userId string, productId string, newExpiresAt int64) error {
	pool := config.GetPool()
	ctx := context.Background()

	subs, _ := GetSubscriptions()

	var tier int
	var subId string

	for _, s := range subs {
		if s.ProductId == productId {
			tier = s.Tier
			subId = s.ID
			break
		}
	}

	update := map[string]any{
		"product_id":        productId,
		"subscription_id":   subId,
		"subscription_tier": tier,
		"expires_at":        newExpiresAt,
		"changed_plan":      false,
	}

	updateJSON, _ := json.Marshal(update)

	_, err := pool.Exec(ctx,
		`UPDATE users SET subscription = subscription || $1::jsonb WHERE id = $2`,
		updateJSON, userId,
	)

	return err
}
