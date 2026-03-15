package chat

// ─── Request types ────────────────────────────────────────────────────────────
// These mirror the TypeScript ChatRequest / PageContext / InteractiveElement types.
// JSON tags must match exactly what the extension sends.

type InteractiveElement struct {
	Type        string `json:"type"`
	Selector    string `json:"selector"`
	Text        string `json:"text,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Value       string `json:"value,omitempty"`
	Href        string `json:"href,omitempty"`
	Name        string `json:"name,omitempty"`
	AriaLabel   string `json:"ariaLabel,omitempty"`
	ID          string `json:"id,omitempty"`
}

type PageContext struct {
	URL                 string               `json:"url"`
	Title               string               `json:"title"`
	PageText            string               `json:"pageText"`
	InteractiveElements []InteractiveElement `json:"interactiveElements"`
}

type ChatMessage struct {
	Role    string `json:"role"` // "user" | "assistant"
	Content string `json:"content"`
}

type ChatRequest struct {
	Message        string      `json:"message"`
	ConversationID string      `json:"conversation_id"`
	PageContext    PageContext  `json:"page_context"`
	History        []ChatMessage `json:"history"`
}

// ─── Response types ───────────────────────────────────────────────────────────

// BrowserStep is a single action inside an execute_steps call.
type BrowserStep struct {
	Action   string   `json:"action"`
	Selector *string  `json:"selector,omitempty"`
	Text     *string  `json:"text,omitempty"`
	Value    *string  `json:"value,omitempty"`
	URL      *string  `json:"url,omitempty"`
	Key      *string  `json:"key,omitempty"`
	X        *float64 `json:"x,omitempty"`
	Y        *float64 `json:"y,omitempty"`
}

type ActionParams struct {
	Selector *string       `json:"selector,omitempty"`
	Text     *string       `json:"text,omitempty"`
	Value    *string       `json:"value,omitempty"`
	URL      *string       `json:"url,omitempty"`
	Key      *string       `json:"key,omitempty"`
	X        *float64      `json:"x,omitempty"`
	Y        *float64      `json:"y,omitempty"`
	Steps    []BrowserStep `json:"steps,omitempty"`
}

type BrowserAction struct {
	Type   string        `json:"type"`
	Params *ActionParams `json:"params,omitempty"`
}

type ChatResponse struct {
	Message        string         `json:"message"`
	Action         *BrowserAction `json:"action,omitempty"`
	ConversationID string         `json:"conversation_id"`
}
