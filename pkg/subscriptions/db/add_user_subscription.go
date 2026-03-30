package subscriptions

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func AddUserSubscription(userId string, productId string, expiresAt int64, platform, purchaseToken, subProductId string) error {
	pool := config.GetPool()
	ctx := context.Background()

	isAnual := strings.HasSuffix(productId, "y")

	lookupProductId := productId
	if isAnual {
		lookupProductId = strings.TrimSuffix(productId, "y")
	}

	sub, err := GetSubscriptionByProductId(lookupProductId)
	if err != nil || sub == nil {
		log.Printf("Subscription not found for productID: %s\n", productId)
		if err != nil {
			return err
		}
		return nil
	}

	userSubscription := authDomain.UserSubscription{
		ProductId:        productId,
		PurchaseToken:    purchaseToken,
		SubProductId:     subProductId,
		SubscriptionId:   sub.ID,
		SubscriptionTier: sub.Tier,
		ExpiresAt:        expiresAt,
		Platform:         platform,
	}

	subscriptionJSON, _ := json.Marshal(userSubscription)

	_, err = pool.Exec(ctx,
		`UPDATE users SET subscription = $1 WHERE id = $2`,
		subscriptionJSON, userId,
	)

	return err
}
