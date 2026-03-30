package subscriptions

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
)

func GetSubscriptionByProductId(productId string) (*subscriptionsDomain.Subscription, error) {
	pool := config.GetPool()
	ctx := context.Background()

	var s subscriptionsDomain.Subscription
	var benefitsJSON []byte

	err := pool.QueryRow(ctx,
		`SELECT id, name, tier, product_id, platform, benefits FROM subscriptions WHERE product_id = $1`,
		productId,
	).Scan(&s.ID, &s.Name, &s.Tier, &s.ProductId, &s.Platform, &benefitsJSON)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if len(benefitsJSON) > 0 {
		json.Unmarshal(benefitsJSON, &s.Benefits)
	}

	return &s, nil
}
