package chat

type Chat struct {
	ID          string        `json:"id"`
	Participants []Participant `json:"participants"`
	Messages    []Message     `json:"messages"`
	Owner       string        `json:"owner"`
	Agent       *ChatAgent    `json:"chatAgent"`
	Deactivated bool          `json:"deactivated"`
	LastMessage *Message      `json:"lastMessage,omitempty"`
}
