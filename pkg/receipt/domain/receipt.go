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

// Request payload for saving receipt
type SaveReceiptParams struct {
	PackageName   string `json:"package_name"`
	ProductID     string `json:"product_id"`
	PurchaseToken string `json:"purchase_token"`
	Wallet        string `json:"wallet"`
	Platform      string `json:"platform"`
}
