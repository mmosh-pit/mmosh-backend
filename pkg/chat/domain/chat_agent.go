package chat

type ChatAgent struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Desc            string `json:"desc"`
	Image           string `json:"image"`
	Symbol          string `json:"symbol"`
	Key             string `json:"key"`
	SystemPrompt    string `json:"system_prompt"`
	CreatorUsername string `json:"creatorUsername"`
	Type            string `json:"type"`
	DefaultModel    string `json:"defaultmodel"`
}
