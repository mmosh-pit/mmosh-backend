package receiptDomain

import "time"

type Receipt struct {
	ID            string    `json:"id,omitempty"`
	PackageName   string    `json:"package_name"`
	ProductID     string    `json:"product_id"`
	PurchaseToken string    `json:"purchase_token"`
	Wallet        string    `json:"wallet"`
	Platform      string    `json:"platform"`
	CreatedAt     time.Time `json:"created_at"`
	ExpiredAt     time.Time `json:"expired_at"`
	IsCanceled    bool      `json:"is_canceled"`
}

type SaveReceiptParams struct {
	PackageName   string `json:"package_name"`
	ProductID     string `json:"product_id"`
	PurchaseToken string `json:"purchase_token"`
	Wallet        string `json:"wallet"`
	Platform      string `json:"platform"`
}

type VerifyReceiptParams struct {
	PurchaseToken string `json:"purchase_token"`
	Wallet        string `json:"wallet"`
}

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
	PaymentStatusCancelled PaymentStatus = "cancelled"
	PaymentStatusUnknown   PaymentStatus = "unknown"
)

type PaymentInfo struct {
	Status           PaymentStatus
	PaymentState     int64
	AutoRenewing     bool
	ExpiryTime       time.Time
	PurchaseTime     time.Time
	CancellationTime *time.Time
	RefundTime       *time.Time
	Message          string
}

type DistributeRequest struct {
	StakedAmount int    `json:"stakedAmount"`
	UserAddress  string `json:"userAddeess"`
	PurchaseID   string `json:"purchaseId"`
}
