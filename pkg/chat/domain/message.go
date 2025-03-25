package chat

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID           *primitive.ObjectID `bson:"_id" json:"id"`
	Content      string              `bson:"content" json:"content"`
	Type         string              `bson:"type" json:"type"`
	CreatedAt    time.Time           `bson:"created_at" json:"created_at"`
	Sender       *primitive.ObjectID `bsob:"sender" json:"sender"`
	IsLoading    bool                `json:"is_loading"`
	SystemPrompt string
	Namespaces   []string
	AgentId      *primitive.ObjectID `json:"agent_id"`
	ChatId       *primitive.ObjectID `json:"chat_id"`
}
