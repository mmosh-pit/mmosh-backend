package domain

type UpdateBotPayload struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	DefaultModel string `json:"defaultmodel"`
	Deactivated  bool   `json:"deactivated"`
}
