package subscriptions

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscription struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Tier      int                `bson:"tier" json:"tier"`
	ProductId string             `bson:"product_id" json:"product_id"`
	Agents    string             `bson:"agents" json:"agents"`
	Platform  string             `bson:"platform" json:"platform"`
}
