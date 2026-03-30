package chat

import "time"

type Message struct {
	ID           string    `json:"id"`
	Content      string    `json:"content"`
	Type         string    `json:"type"`
	CreatedAt    time.Time `json:"created_at"`
	Sender       string    `json:"sender"`
	IsLoading    bool      `json:"is_loading"`
	SystemPrompt string    `json:"-"`
	Namespaces   []string  `json:"-"`
	AgentId      string    `json:"agent_id"`
	ChatId       string    `json:"chat_id"`
}
