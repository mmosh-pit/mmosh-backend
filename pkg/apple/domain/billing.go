package apple

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var (
	// The notificationType for a in-app purchase or external purchase can have the following values:
	SUBSCRIBED                string = "SUBSCRIBED"
	DID_CHANGE_RENEWAL_PREF   string = "DID_CHANGE_RENEWAL_PREF"
	DID_CHANGE_RENEWAL_STATUS string = "DID_CHANGE_RENEWAL_STATUS"
	OFFER_REDEEMED            string = "OFFER_REDEEMED"
	DID_RENEW                 string = "DID_RENEW"
	EXPIRED                   string = "EXPIRED"
	DID_FAIL_TO_RENEW         string = "DID_FAIL_TO_RENEW"
	GRACE_PERIOD_EXPIRED      string = "GRACE_PERIOD_EXPIRED"
	PRICE_INCREASE            string = "PRICE_INCREASE"
	REFUND                    string = "REFUND"
	REFUND_DECLINED           string = "REFUND_DECLINED"
	CONSUMPTION_REQUEST       string = "CONSUMPTION_REQUEST"
	RENEWAL_EXTENDED          string = "RENEWAL_EXTENDED"
	REVOKE                    string = "REVOKE"
	TEST                      string = "TEST"
	RENEWAL_EXTENSION         string = "RENEWAL_EXTENSION"
	REFUND_REVERSED           string = "REFUND_REVERSED"
	EXTERNAL_PURCHASE_TOKEN   string = "EXTERNAL_PURCHASE_TOKEN"
	ONE_TIME_CHARGE           string = "ONE_TIME_CHARGE"
	// End of values

	// The subtype for a in-app purchase or external purchase can have the following values:
	INITIAL_BUY         string = "INITIAL_BUY"
	RESUBSCRIBE         string = "RESUBSCRIBE"
	DOWNGRADE           string = "DOWNGRADE"
	UPGRADE             string = "UPGRADE"
	AUTO_RENEW_ENABLED  string = "AUTO_RENEW_ENABLED"
	AUTO_RENEW_DISABLED string = "AUTO_RENEW_DISABLED"
	VOLUNTARY           string = "VOLUNTARY"
	BILLING_RETRY       string = "BILLING_RETRY"
	// -- PRICE_INCREASE
	GRACE_PERIOD         string = "GRACE_PERIOD"
	PENDING              string = "PENDING"
	ACCEPTED             string = "ACCEPTED"
	BILLING_RECOVERY     string = "BILLING_RECOVERY"
	PRODUCT_NOT_FOR_SALE string = "PRODUCT_NOT_FOR_SALE"
	SUMMARY              string = "SUMMARY"
	FAILURE              string = "FAILURE"
	UNREPORTED           string = "UNREPORTED"
	// End of values

	// The consumptionRequestReason can have the following values:
	UNINTENDED_PURCHASE string = "UNINTENDED_PURCHASE"
	FULFILLMENT_ISSUE   string = "FULFILLMENT_ISSUE"
	LEGAL               string = "LEGAL"
	OTHER               string = "OTHER"
	// End of values

	// The environment can have the following values:
	Sandbox    string = "Sandbox"
	Production string = "Production"
	// End of values

	// The autoRenewStatus for JWSRenewalInfo can have the following values:
	AUTOMATIC_RENEWAL_IS_OFF int32 = 0
	AUTOMATIC_RENEWAL_IS_ON  int32 = 1
	// End of values

	// The expirationIntent for JWSRenewalInfo can have the following values:
	EXPIRATION_INTENT_CANCELED_BY_USER                                       int32 = 1
	EXPIRATION_INTENT_CANCELED_BY_BILLING_ERROR                              int32 = 2
	EXPIRATION_INTENT_CANCELED_BY_NO_CONSENT_FOR_PRIE_INCREASE               int32 = 3
	EXPIRATION_INTENT_CANCELED_BY_PRODUCT_NOT_AVAILABLE                      int32 = 4
	EXPIRATION_INTENT_CANCELED_BY_SUBSCRIPTION_EXPIRED_FOR_SOME_OTHER_REASON int32 = 5
	// End of values

	// The offerDiscountType for JWSRenewalInfo and JWSTransaction can have the following values:
	FREE_TRIAL    string = "FREE_TRIAL"
	PAY_AS_YOU_GO string = "PAY_AS_YOU_GO"
	PAY_UP_FRONT  string = "PAY_UP_FRONT"
	// End of values

	// The offerType for JWSRenewalInfo and JWSTransaction can have the following values:
	INTRODUCTORY_OFFER      int32 = 1
	PROMOTIONAL_OFFER       int32 = 2
	SUBSCRIPTION_CODE_OFFER int32 = 3
	WIN_BACK_OFFER          int32 = 4
	// End of values

	// The priceIncreaseStatus for JWSRenewalInfo can have the following values:
	NO_RESPONSE_FOR_AUTO_RENEWABLE_SUBSCRIPTION_PRICE_INCREASE    int32 = 0
	USER_CONSENTED_FOR_AUTO_RENEWABLE_SUBSCRIPTION_PRICE_INCREASE int32 = 1
	// End of values

	// The priceIncreaseStatus for JWSRenewalInfo and inAppOwnershipType for JWSTransaction can have the following values:
	FAMILY_SHARED string = "FAMILY_SHARED"
	PURCHASED     string = "PURCHASED"
	// End of values

	// The inAppOwnershipType for JWSTransaction can have the following values:
	AUTO_RENEWABLE_SUBSCRIPTION_IS_ACTIVE               int32 = 1
	AUTO_RENEWABLE_SUBSCRIPTION_IS_EXPIRED              int32 = 2
	AUTO_RENEWABLE_SUBSCRIPTION_IN_BILLING_RETRY_PERIOD int32 = 3
	AUTO_RENEWABLE_SUBSCRIPTION_IN_BILLING_GRACE_PERIOD int32 = 4
	AUTO_RENEWABLE_SUBSCRIPTION_IS_REVOKED              int32 = 5
	// End of values

	// The revocationReason for JWSTransaction can have the following values:
	REVOCATION_BY_USER_FOR_OTHER_REASONS int32 = 0
	REVOCATION_BY_USER_FOR_APP_ISSUE     int32 = 1
	// End of values

	// The transactionReason for JWSTransaction can have the following values:
	PURCHASE string = "PURCHASE"
	RENEWAL  string = "RENEWAL"
	// End of values

	// The type for JWSTransaction can have the following values:
	AUTO_RENEWABLE_SUBSCRIPTION string = "Auto-Renewable Subscription"
	NON_CONSUMABLE              string = "Non-Consumable"
	CONSUMABLE                  string = "Consumable"
	NON_RENEWING_SUBSCRIPTION   string = "Non-Renewing Subscription"
	// End of values
)

type JWSRenewalInfo struct {
	AutoRenewProductId          string    `json:"autoRenewProductId"`      // The product identifier of the product that renews at the next billing period.
	AutoRenewStatus             int32     `json:"autoRenewStatus"`         // The renewal status of the auto-renewable subscription.
	Currency                    string    `json:"currency"`                // The three-letter ISO 4217 currency code for the price of the product.
	EligibleWinBackOfferIds     []string  `json:"eligibleWinBackOfferIds"` // An array of win-back offer identifiers that a customer is eligible to redeem, which sorts the identifiers to present the better offers first.
	Environment                 string    `json:"environment"`
	ExpirationIntent            int32     `json:"expirationIntent"`
	GracePeriodExpiresDate      time.Time `json:"gracePeriodExpiresDate"`
	IsInBillingRetryPeriod      bool      `json:"isInBillingRetryPeriod"`
	OfferDiscountType           string    `json:"offerDiscountType"`
	OfferIdentifier             string    `json:"offerIdentifier"` // The string identifier of a subscription offer that you create in App Store Connect.
	OfferType                   string    `json:"offerType"`
	OriginalTransactionId       string    `json:"originalTransactionId"` // The original transaction identifier of a purchase.
	PriceIncreaseStatus         int32     `json:"priceIncreaseStatus"`
	ProductId                   string    `json:"productId"` // The product identifier of the In-App Purchase.
	RecentSubscriptionStartDate time.Time `json:"recentSubscriptionStartDate"`
	RenewalDate                 int64     `json:"renewalDate"`  // The UNIX time, in milliseconds, when the most recent auto-renewable subscription purchase expires.
	RenewalPrice                int64     `json:"renewalPrice"` // The renewal price, in milliunits, of the auto-renewable subscription that renews at the next billing period.
	SignedDate                  int64     `json:"signedDate"`
}

type JWSTransaction struct {
	AppAccountToken             string `json:"appAccountToken"` // A UUID you create at the time of purchase that associates the transaction with a customer on your own service. If your app doesn’t provide an appAccountToken, this string is empty. For more information, see appAccountToken(_:).
	BundleId                    string `json:"bundleId"`        // The bundle identifier of the app.
	Currency                    string `json:"currency"`        // The three-letter ISO 4217 currency code for the price of the product.
	Environment                 string `json:"environment"`
	ExpiresDate                 int64  `json:"expiresDate"` // The UNIX time, in milliseconds, that the subscription expires or renews.
	InAppOwnershipType          string `json:"inAppOwnershipType"`
	IsUpgraded                  bool   `json:"isUpgraded"` // A Boolean value that indicates whether the customer upgraded to another subscription.
	OfferDiscountType           string `json:"offerDiscountType"`
	OfferIdentifier             string `json:"offerIdentifier"` // The string identifier of a subscription offer that you create in App Store Connect.
	OfferType                   string `json:"offerType"`
	OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`  // The purchase date of the transaction associated with the original transaction identifier.
	OriginalTransactionId       string `json:"originalTransactionId"` // The original transaction identifier of a purchase.
	Price                       int64  `json:"price"`                 // The price, in milliunits, of the In-App Purchase that the system records in the transaction.
	ProductId                   string `json:"productId"`             // The product identifier of the In-App Purchase.
	PurchaseDate                int64  `json:"purchaseDate"`          // The UNIX time, in milliseconds, that the App Store charged the user’s account for a purchase, restored product, subscription, or subscription renewal after a lapse.
	Quantity                    int32  `json:"quantity"`              // The number of consumable products the user purchased.
	RevocationDate              int64  `json:"revocationDate"`        // The UNIX time, in milliseconds, that the App Store refunded the transaction or revoked it from Family Sharing.
	RevocationReason            string `json:"revocationReason"`
	SignedDate                  int64  `json:"signedDate"`
	Storefront                  string `json:"storefront"`                  // The three-letter code that represents the country or region associated with the App Store storefront for the purchase.
	StorefrontId                string `json:"storefrontId"`                // An Apple-defined value that uniquely identifies the App Store storefront associated with the purchase.
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"` // The identifier of the subscription group that the subscription belongs to.
	TransactionId               string `json:"transactionId"`               // The unique identifier of the transaction. The App Store generates a new value for transaction identifier every time the subscription automatically renews or the user restores it on a new device.
	TransactionReason           string `json:"transactionReason"`
	Type                        string `json:"type"`               // The product type of the In-App Purchase.
	WebOrderLineItemId          string `json:"webOrderLineItemId"` // The unique identifier of subscription purchase events across devices, including subscription renewals.
}

func (J JWSTransaction) Valid() error {
	return nil
}

type DecodedPayloadData struct {
	AppAppleId               int64  `json:"appAppleId"`               // The unique identifier of the app. This property is available for apps that users download from the App Store. It isn’t present in the sandbox environment.
	BundleId                 string `json:"bundleId"`                 // The bundle identifier of the app.
	BundleVersion            string `json:"bundleVersion"`            // The version of the build that identifies an iteration of the bundle.
	ConsumptionRequestReason string `json:"consumptionRequestReason"` // The reason the customer requested the refund. This field appears only for CONSUMPTION_REQUEST notifications, which the server sends when a customer initiates a refund request for a consumable in-app purchase or auto-renewable subscription.
	Environment              string `json:"environment"`
	SignedRenewalInfo        string `json:"signedRenewalInfo"`     // Subscription renewal information signed by the App Store, in JSON Web Signature (JWS) format. This field appears only for notifications that apply to auto-renewable subscriptions.
	SignedTransactionInfo    string `json:"signedTransactionInfo"` // Transaction information signed by the App Store, in JSON Web Signature (JWS) format.
	Status                   int32  `json:"status"`                // The status of an auto-renewable subscription as of the signedDate in the responseBodyV2DecodedPayload. This field appears only for notifications sent for auto-renewable subscriptions.
}

// The payload data for a subscription-renewal-date extension notification.
type DecodedPayloadSummary struct {
	RequestIdentifier      uuid.UUID `json:"requestIdentifier"` // The UUID that represents a specific request to extend a subscription renewal date. This value matches the value you initially specify in the requestIdentifier when you call Extend Subscription Renewal Dates for All Active Subscribers in the App Store Server API.
	Environment            string    `json:"environment"`
	AppAppleId             int64     `json:"appAppleId"`             // The unique identifier of the app. This property is available for apps that users download from the App Store. It isn’t present in the sandbox environment.
	BundleId               string    `json:"bundleId"`               // The bundle identifier of the app.
	ProductId              string    `json:"productId"`              // The product identifier of the auto-renewable subscription that the subscription-renewal-date extension applies to.
	StorefrontCountryCodes []string  `json:"storefrontCountryCodes"` // A list of country codes that limits the App Store’s attempt to apply the subscription-renewal-date extension. If this list isn’t present, the subscription-renewal-date extension applies to all storefronts.
	FailedCount            int64     `json:"failedCount"`            // The final count of subscriptions that fail to receive a subscription-renewal-date extension.
	SucceededCount         int64     `json:"succeededCount"`         // The final count of subscriptions that successfully receive a subscription-renewal-date extension.
}

type DecodedPayloadExternalPurchaseToken struct {
	ExternalPurchaseId string `json:"externalPurchaseId"` // The unique identifier of the token. Use this value to report tokens and their associated transactions in the Send External Purchase Report endpoint.
	TokenCreationDate  int64  `json:"tokenCreationDate"`
	AppAppleId         int64  `json:"appAppleId"` // The app Apple ID for which the system generated the token.
	BundleId           string `json:"bundleId"`   // The bundle ID of the app for which the system generated the token.
}

// The data, summary, and externalPurchaseToken fields are mutually exclusive. The payload contains only one of these fields.
type ResponseBodyV2DecodedPayload struct {
	jwt.RegisteredClaims
	NotificationType      string                              `json:"notificationType"`                // The in-app purchase event for which the App Store sends this version 2 notification.
	SubType               string                              `json:"subtype"`                         // Additional information that identifies the notification event. The subtype field is present only for specific version 2 notifications.
	Data                  DecodedPayloadData                  `json:"data,omitempty"`                  // The object that contains the app metadata and signed renewal and transaction information.
	Summary               DecodedPayloadSummary               `json:"summary,omitempty"`               // The summary data that appears when the App Store server completes your request to extend a subscription renewal date for eligible subscribers. For more information, see Extend Subscription Renewal Dates for All Active Subscribers.
	ExternalPurchaseToken DecodedPayloadExternalPurchaseToken `json:"externalPurchaseToken,omitempty"` // This field appears when the notificationType is EXTERNAL_PURCHASE_TOKEN.
	Version               string                              `json:"version"`
	SignedDate            int64                               `json:"signedDate"`
	NotificationUUID      string                              `json:"notificationUUID"`
}

type ResponseBodyV2 struct {
	SignedPayload string `json:"signedPayload"`
}
