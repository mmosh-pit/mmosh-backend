package subscriptions

type Subscription struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Tier      int      `json:"tier"`
	ProductId string   `json:"product_id"`
	Platform  string   `json:"platform"`
	Benefits  []string `json:"benefits"`
}
