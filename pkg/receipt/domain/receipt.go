package receiptDomain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receipt struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PackageName   string             `bson:"package_name" json:"package_name"`
	ProductID     string             `bson:"product_id" json:"product_id"`
	PurchaseToken string             `bson:"purchase_token" json:"purchase_token"`
	Wallet        string             `bson:"wallet" json:"wallet"`
	Platform      string             `bson:"platform" json:"platform"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
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
