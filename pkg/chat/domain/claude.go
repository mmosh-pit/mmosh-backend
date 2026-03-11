package chat

// ─── Request types ────────────────────────────────────────────────────────────
// These mirror the TypeScript ChatRequest / PageContext / TweetInfo types.
// JSON tags must match exactly what the extension sends.

type TweetInfo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
	Handle string `json:"handle"`
	URL    string `json:"url"`
}

type PageContext struct {
	URL      string      `json:"url"`
	PageType string      `json:"pageType"` // camelCase — matches TypeScript PageContext field
	Tweets   []TweetInfo `json:"tweets"`
}

type ChatMessage struct {
	Role    string `json:"role"` // "user" | "assistant"
	Content string `json:"content"`
}

type ChatRequest struct {
	Message        string        `json:"message"`
	ConversationID string        `json:"conversation_id"`
	PageContext    PageContext   `json:"page_context"`
	History        []ChatMessage `json:"history"`
}

// ─── Response types ───────────────────────────────────────────────────────────
// ChatResponse is returned directly (no data wrapper).

type ActionParams struct {
	Text     *string `json:"text,omitempty"`
	TweetURL *string `json:"tweet_url,omitempty"`
	URL      *string `json:"url,omitempty"`
}

type TwitterAction struct {
	Type   string        `json:"type"`
	Params *ActionParams `json:"params,omitempty"`
}

type ChatResponse struct {
	Message        string         `json:"message"`
	Action         *TwitterAction `json:"action,omitempty"`
	ConversationID string         `json:"conversation_id"`
}
