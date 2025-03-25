package google

var (
	// The notificationType for a subscription can have the following values:
	SUBSCRIPTION_RECOVERED                 int = 1
	SUBSCRIPTION_RENEWED                   int = 2
	SUBSCRIPTION_CANCELED                  int = 3
	SUBSCRIPTION_PURCHASED                 int = 4
	SUBSCRIPTION_ON_HOLD                   int = 5
	SUBSCRIPTION_IN_GRACE_PERIOD           int = 6
	SUBSCRIPTION_RESTARTED                 int = 7
	SUBSCRIPTION_PRICE_CHANGE_CONFIRMED    int = 8
	SUBSCRIPTION_DEFERRED                  int = 9
	SUBSCRIPTION_PAUSED                    int = 10
	SUBSCRIPTION_PAUSE_SCHEDULE_CHANGED    int = 11
	SUBSCRIPTION_REVOKED                   int = 12
	SUBSCRIPTION_EXPIRED                   int = 13
	SUBSCRIPTION_PENDING_PURCHASE_CANCELED int = 20
	// End of values

	SUBSCRIPTION_STATE_UNSPECIFIED               string = "SUBSCRIPTION_STATE_UNSPECIFIED"
	SUBSCRIPTION_STATE_PENDING                   string = "SUBSCRIPTION_STATE_PENDING"
	SUBSCRIPTION_STATE_ACTIVE                    string = "SUBSCRIPTION_STATE_ACTIVE"
	SUBSCRIPTION_STATE_PAUSED                    string = "SUBSCRIPTION_STATE_PAUSED"
	SUBSCRIPTION_STATE_IN_GRACE_PERIOD           string = "SUBSCRIPTION_STATE_IN_GRACE_PERIOD"
	SUBSCRIPTION_STATE_ON_HOLD                   string = "SUBSCRIPTION_STATE_ON_HOLD"
	SUBSCRIPTION_STATE_CANCELED                  string = "SUBSCRIPTION_STATE_CANCELED"
	SUBSCRIPTION_STATE_EXPIRED                   string = "SUBSCRIPTION_STATE_EXPIRED"
	SUBSCRIPTION_STATE_PENDING_PURCHASE_CANCELED string = "SUBSCRIPTION_STATE_PENDING_PURCHASE_CANCELED"

	SUBSCRIPTION_CANCEL_SURVEY_REASON_UNSPECIFIED      string = "CANCEL_SURVEY_REASON_UNSPECIFIED"
	SUBSCRIPTION_CANCEL_SURVEY_REASON_NOT_ENOUGH_USAGE string = "CANCEL_SURVEY_REASON_NOT_ENOUGH_USAGE"
	SUBSCRIPTION_CANCEL_SURVEY_REASON_TECHNICAL_ISSUES string = "CANCEL_SURVEY_REASON_TECHNICAL_ISSUES"
	SUBSCRIPTION_CANCEL_SURVEY_REASON_COST_RELATED     string = "CANCEL_SURVEY_REASON_COST_RELATED"
	SUBSCRIPTION_CANCEL_SURVEY_REASON_FOUND_BETTER_APP string = "CANCEL_SURVEY_REASON_FOUND_BETTER_APP"
	SUBSCRIPTION_CANCEL_SURVEY_REASON_OTHERS           string = "CANCEL_SURVEY_REASON_OTHERS"

	SUBSCRIPTION_ACKNOWLEDGEMENT_STATE_UNSPECIFIED  string = "ACKNOWLEDGEMENT_STATE_UNSPECIFIED"
	SUBSCRIPTION_ACKNOWLEDGEMENT_STATE_PENDING      string = "ACKNOWLEDGEMENT_STATE_PENDING"
	SUBSCRIPTION_ACKNOWLEDGEMENT_STATE_ACKNOWLEDGED string = "ACKNOWLEDGEMENT_STATE_ACKNOWLEDGED"

	PRICE_CHANGE_MODE_UNSPECIFIED string = "PRICE_CHANGE_MODE_UNSPECIFIED"
	PRICE_DECREASE                string = "PRICE_DECREASE"
	PRICE_INCREASE                string = "PRICE_INCREASE"
	OPT_OUT_PRICE_INCREASE        string = "OPT_OUT_PRICE_INCREASE"

	PRICE_CHANGE_STATE_UNSPECIFIED string = "PRICE_CHANGE_STATE_UNSPECIFIED"
	OUTSTANDING                    string = "OUTSTANDING"
	CONFIRMED                      string = "CONFIRMED"
	APPLIED                        string = "APPLIED"

	// The notificationType for a product can have the following values:
	ONE_TIME_PRODUCT_PURCHASED int = 1
	ONE_TIME_PRODUCT_CANCELED  int = 2
	// End of values

	PURCHASE_STATE_PURCHASED int64 = 0
	PURCHASE_STATE_CANCELED  int64 = 1
	PURCHASE_STATE_PENDING   int64 = 2

	CONSUMPTION_STATE_YET_TO_BE_CONSUMED int64 = 0
	CONSUMPTION_STATE_CONSUMED           int64 = 1

	PURCHASE_TYPE_TEST     int = 0
	PURCHASE_TYPE_PROMO    int = 1
	PURCHASE_TYPE_REWARDED int = 2

	ACKNOWLEDGEMENT_STATE_YET_TO_BE_ACKNOWLEDGED int = 0
	ACKNOWLEDGEMENT_STATE_ACKNOWLEDGED           int = 1

	// The productType for a voided purchase can have the following values:
	PRODUCT_TYPE_SUBSCRIPTION int = 1
	PRODUCT_TYPE_ONE_TIME     int = 2
	// End of values

	// The refundType for a voided purchase can have the following values:
	REFUND_TYPE_FULL_REFUND                   int = 1
	REFUND_TYPE_QUANTITY_BASED_PARTIAL_REFUND int = 2
	// End of values
)

type Money struct {
	CurrencyCode string `json:"currencyCode"`
	Units        string `json:"units"`
	Nanos        int    `json:"nanos"`
}

type SubscriptionNotification struct {
	Version          string `json:"version"` // The version of this notification
	NotificationType int    `json:"notificationType"`
	PurchaseToken    string `json:"purchaseToken"`  // The token provided to the user's device when the product was purchased
	SubscriptionId   string `json:"subscriptionId"` // The purchased subscription's product ID (for example, "monthly001")
}

type OneTimeProductNotification struct {
	Version          string `json:"version"` // The version of this notification
	NotificationType int    `json:"notificationType"`
	PurchaseToken    string `json:"purchaseToken"` // The token provided to the user's device when the product was purchased
	Sku              string `json:"sku"`           // The purchased one-time product ID (for example, "sword_001")
}

type VoidedPurchaseNotification struct {
	PurchaseToken string `json:"purchaseToken"` // The token provided to the user's device when the product was purchased
	OrderId       string `json:"orderId"`       // This identifies the order associated with the voided transaction
	ProductType   int    `json:"productType"`
	RefundType    int    `json:"refundType"`
}

type TestNotification struct {
	Version string `json:"version"` // The version of this notification
}

type MessageData struct {
	Version                    string                     `json:"version"`     // The version of this notification
	PackageName                string                     `json:"packageName"` // This identifies the app
	EventTimeMillis            string                     `json:"eventTimeMillis"`
	SubscriptionNotification   SubscriptionNotification   `json:"subscriptionNotification,omitempty"`
	OneTimeProductNotification OneTimeProductNotification `json:"oneTimeProductNotification,omitempty"`
	VoidedPurchaseNotification VoidedPurchaseNotification `json:"voidedPurchaseNotification,omitempty"`
	TestNotification           TestNotification           `json:"testNotification,omitempty"`
}

type PushSubscriptionMessage struct {
	Attributes  map[string]interface{} `json:"attributes"`
	Data        string                 `json:"data"`
	MessageId   string                 `json:"message_id"`
	PublishTime string                 `json:"publish_time"`
}

type PushSubscription struct {
	DeliveryAttempt int                     `json:"deliveryAttempt"`
	Message         PushSubscriptionMessage `json:"message"`
	Subscription    string                  `json:"subscription"`
}
