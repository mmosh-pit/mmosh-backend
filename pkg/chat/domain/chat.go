package chat

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID           *primitive.ObjectID `bson:"_id" json:"id"`
	Participants []Participant       `bson:"participants" json:"participants"`
	Messages     []Message           `bson:"messages" json:"messages"`
	Owner        *primitive.ObjectID `bson:"owner" json:"owner"`
	Agent        *ChatAgent          `bson:"chatAgent" json:"chatAgent"`
	Deactivated  bool                `bson:"deactivated" json:"deactivated"`
	LastMessage  Message             `bson:"lastMessage" json:"lastMessage"`
}
