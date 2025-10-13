package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	googleDomain "github.com/mmosh-pit/mmosh_backend/pkg/google/domain"
	subscriptionsDb "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/db"
	"google.golang.org/api/androidpublisher/v3"
)

func HandleSubscriptionNotification(packageName string, data googleDomain.SubscriptionNotification) error {

	log.Printf("Got a notification here! %v, %v, %v\n", data.SubscriptionId, data.Version, data.NotificationType)

	if data.NotificationType == googleDomain.SUBSCRIPTION_PURCHASED || data.NotificationType == googleDomain.SUBSCRIPTION_RENEWED ||
		data.NotificationType == googleDomain.SUBSCRIPTION_DEFERRED || data.NotificationType == googleDomain.SUBSCRIPTION_RECOVERED ||
		data.NotificationType == googleDomain.SUBSCRIPTION_EXPIRED || data.NotificationType == googleDomain.SUBSCRIPTION_ON_HOLD ||
		data.NotificationType == googleDomain.SUBSCRIPTION_PENDING_PURCHASE_CANCELED || data.NotificationType == googleDomain.SUBSCRIPTION_CANCELED ||
		data.NotificationType == googleDomain.SUBSCRIPTION_RESTARTED {
		apService, err := androidpublisher.NewService(context.Background())
		if err != nil {
			log.Println("androidpublisher.NewService: ", err)
			return err
		}

		subscriptionPurchase, err := apService.Purchases.Subscriptionsv2.Get(packageName, data.PurchaseToken).Do()
		if err != nil {
			log.Println("apService.Purchases.Subscriptionsv2.Get: ", err)
			return err
		}

		id, err := uuid.Parse(subscriptionPurchase.ExternalAccountIdentifiers.ObfuscatedExternalAccountId)
		if err != nil {
			return err
		}

		user, err := auth.GetUserByUuidId(id.String())

		if err != nil {
			return err
		}

		// TEMPORAL
		// jsonData, err := json.MarshalIndent(subscriptionPurchase, "", "  ") // Use MarshalIndent for pretty-printed JSON
		// if err == nil {
		// 	// Write the JSON data to a file
		// 	filePath := "output_google.json"
		// 	err = os.WriteFile(filePath, jsonData, 0644) // 0644 sets file permissions
		// 	if err != nil {
		// 		fmt.Printf("Error writing to file: %v\n", err)
		// 	}
		// }
		// END TEMPORAL

		parsedObjectId := user.ID

		log.Printf("AcknowledgementState: %v\n", subscriptionPurchase.AcknowledgementState)
		log.Printf("SubscriptionState: %v\n", subscriptionPurchase.SubscriptionState)
		productId := ""
		subProductId := ""

		for _, spli := range subscriptionPurchase.LineItems {
			subProductId = spli.OfferDetails.BasePlanId
		}

		if productId == "" {
			productId = data.SubscriptionId
		}

		if data.NotificationType == googleDomain.SUBSCRIPTION_CANCELED && subscriptionPurchase.CanceledStateContext != nil &&
			subscriptionPurchase.CanceledStateContext.UserInitiatedCancellation != nil {

			subscriptionsDb.DeleteUserSubscription(parsedObjectId, productId)

		} else if data.NotificationType == googleDomain.SUBSCRIPTION_RESTARTED {
		} else if subscriptionPurchase.SubscriptionState == googleDomain.SUBSCRIPTION_STATE_ACTIVE {
			if subscriptionPurchase.AcknowledgementState == googleDomain.SUBSCRIPTION_ACKNOWLEDGEMENT_STATE_PENDING {
				if err := apService.Purchases.Subscriptions.Acknowledge(
					packageName,
					data.SubscriptionId,
					data.PurchaseToken,
					&androidpublisher.SubscriptionPurchasesAcknowledgeRequest{
						DeveloperPayload: fmt.Sprintf(`{"developerPayload": "%s:approved"}`, subscriptionPurchase.ExternalAccountIdentifiers.ObfuscatedExternalAccountId),
					},
				).Do(); err != nil {
					log.Println("apService.Purchases.Subscriptions.Acknowledge: ", err)
					return err
				}

				expiresAt, err := time.Parse(time.RFC3339Nano, subscriptionPurchase.LineItems[0].ExpiryTime)

				if err != nil {
					return err
				}

				log.Println("[GOOGLE] Adding user subscription...")
				subscriptionsDb.AddUserSubscription(parsedObjectId, productId, expiresAt.Unix(), "google", data.PurchaseToken, subProductId)

			}

		} else if subscriptionPurchase.SubscriptionState == googleDomain.SUBSCRIPTION_STATE_PENDING_PURCHASE_CANCELED ||
			subscriptionPurchase.SubscriptionState == googleDomain.SUBSCRIPTION_STATE_EXPIRED {

			subscriptionsDb.DeleteUserSubscription(parsedObjectId, productId)

		}
	}

	return nil
}
