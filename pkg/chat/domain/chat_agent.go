package chat

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatAgent struct {
	Id              *primitive.ObjectID `bson:"_id" json:"id"`
	Name            string              `bson:"name" json:"name"`
	Desc            string              `bson:"desc" json:"desc"`
	Image           string              `bson:"image" json:"image"`
	Symbol          string              `bson:"symbol" json:"symbol"`
	Key             string              `bson:"key" json:"key"`
	SystemPrompt    string              `bson:"system_prompt" json:"system_prompt"`
	CreatorUsername string              `bson:"creatorUsername" json:"creatorUsername"`
	Type            string              `bson:"type" json:"type"`
}
