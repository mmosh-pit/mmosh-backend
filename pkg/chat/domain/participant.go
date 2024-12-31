package chat

import "go.mongodb.org/mongo-driver/bson/primitive"

type Participant struct {
	ID      *primitive.ObjectID `bson:"_id" json:"id"`
	Name    string              `bson:"name" json:"name"`
	Type    string              `bson:"type" json:"type"`
	Picture string              `bson:"picture" json:"picture"`
}
