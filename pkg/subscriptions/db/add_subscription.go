package subscriptions

import (
	"context"
	"encoding/json"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
)

func AddSubscription(data *subscriptionsDomain.Subscription) error {
	pool := config.GetPool()
	ctx := context.Background()

	benefitsJSON, _ := json.Marshal(data.Benefits)

	err := pool.QueryRow(ctx,
		`INSERT INTO subscriptions (name, tier, product_id, platform, benefits)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id`,
		data.Name, data.Tier, data.ProductId, data.Platform, benefitsJSON,
	).Scan(&data.ID)

	return err
}
