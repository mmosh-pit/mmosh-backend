package subscriptions

import (
	"context"
	"encoding/json"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
)

func GetSubscriptions() ([]subscriptionsDomain.Subscription, error) {
	pool := config.GetPool()
	ctx := context.Background()

	var subs []subscriptionsDomain.Subscription

	rows, err := pool.Query(ctx,
		`SELECT id, name, tier, product_id, platform, benefits FROM subscriptions ORDER BY tier DESC`,
	)

	if err != nil {
		return subs, err
	}
	defer rows.Close()

	for rows.Next() {
		var s subscriptionsDomain.Subscription
		var benefitsJSON []byte

		if err := rows.Scan(&s.ID, &s.Name, &s.Tier, &s.ProductId, &s.Platform, &benefitsJSON); err != nil {
			return subs, err
		}

		if len(benefitsJSON) > 0 {
			json.Unmarshal(benefitsJSON, &s.Benefits)
		}

		subs = append(subs, s)
	}

	return subs, rows.Err()
}
