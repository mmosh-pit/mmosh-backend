package apple

import (
	"log"

	"github.com/google/uuid"
	appleDomain "github.com/mmosh-pit/mmosh_backend/pkg/apple/domain"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	subscriptionsDb "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/db"
)

func HandleSubscriptionPurchases(notificationType, notificationSubType string, data appleDomain.JWSTransaction) error {

	log.Printf("Notification data: %v\n", data.IsUpgraded)
	log.Printf("Upgraded to: %v\n", data.ProductId)
	log.Printf("Notification type: %v\n", notificationType)

	id, err := uuid.Parse(data.AppAccountToken)
	if err != nil {
		return err
	}

	user, err := auth.GetUserByUuidId(id.String())

	if err != nil {
		return err
	}

	if notificationType == appleDomain.SUBSCRIBED || notificationType == appleDomain.DID_RENEW ||
		notificationType == appleDomain.OFFER_REDEEMED || notificationType == appleDomain.REFUND_REVERSED {

		subscriptionsDb.AddUserSubscription(user.ID, data.ProductId, data.ExpiresDate, "apple")

		// if notificationType == appleDomain.SUBSCRIBED || notificationType == appleDomain.OFFER_REDEEMED {
		// }

	} else if notificationType == appleDomain.EXPIRED || notificationType == appleDomain.REFUND ||
		notificationType == appleDomain.REVOKE {

		subscriptionsDb.DeleteUserSubscription(user.ID, data.ProductId)

		if notificationType == appleDomain.REFUND {
			// Save in transaction history
		}
	} else if notificationType == appleDomain.DID_CHANGE_RENEWAL_STATUS {
		subscriptionsDb.UpdateSubscriptionStatus(user.ID, data.ProductId, data.ExpiresDate)
	} else if notificationType == appleDomain.DID_CHANGE_RENEWAL_PREF {
		subscriptionsDb.AddChangePlanBadgeToSubscription(user.ID)
	}

	return nil
}
