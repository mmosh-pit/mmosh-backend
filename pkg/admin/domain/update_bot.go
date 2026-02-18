package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateBotPayload struct {
	ID           *primitive.ObjectID `bson:"_id,omitempty"`
	Name         string              `bson:"name" json:"name"`
	Symbol       string              `bson:"symbol" json:"symbol"`
	DefaultModel string              `json:"defaultmodel" bson:"defaultmodel"`
	Deactivated  bool                `json:"deactivated" bson:"deactivated"`
}
