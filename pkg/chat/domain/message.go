package chat

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      *primitive.ObjectID `bson:"_id" json:"id"`
	Content string              `bson:"content" json:"content"`
	Type    string              `bson:"type" json:"type"`
	Sender  *primitive.ObjectID `bsob:"sender" json:"sender"`
}
