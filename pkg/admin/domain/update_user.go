package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateUserPayload struct {
	ID          *primitive.ObjectID `bson:"_id,omitempty"`
	Name        string              `bson:"name" json:"name"`
	Email       string              `bson:"email" json:"email"`
	Role        string              `json:"role" bson:"role"`
	Deactivated bool                `json:"deactivated" bson:"deactivated"`
}
