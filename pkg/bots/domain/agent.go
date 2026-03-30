package bots

import "time"

type Bot struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	Desc             string    `json:"desc"`
	Image            string    `json:"image"`
	Symbol           string    `json:"symbol"`
	Key              string    `json:"key"`
	Price            int       `json:"price"`
	PreSaleStartDate string    `json:"presalestartdate"`
	SystemPrompt     string    `json:"system_prompt"`
	CreatorUsername  string    `json:"creatorUsername"`
	Type             string    `json:"type"`
	DefaultModel     string    `json:"defaultmodel"`
	Deactivated      bool      `json:"deactivated"`
	CreatedAt        time.Time `json:"createdAt"`
}

type ToggleActivateAgentData struct {
	AgentId string `json:"agentId"`
}

type CreateAgentData struct {
	Id               string       `json:"id"`
	Name             string       `json:"name"`
	Symbol           string       `json:"symbol"`
	Desc             string       `json:"desc"`
	Image            string       `json:"image"`
	InviteImage      string       `json:"inviteimage"`
	Key              string       `json:"key"`
	Lut              string       `json:"lut"`
	Seniority        int          `json:"seniority"`
	Price            float64      `json:"price"`
	Distribution     Distribution `json:"distribution"`
	InvitationPrice  float64      `json:"invitationprice"`
	Discount         float64      `json:"discount"`
	Telegram         string       `json:"telegram"`
	Twitter          string       `json:"twitter"`
	Website          string       `json:"website"`
	PresaleSupply    int          `json:"presalesupply"`
	MinPresaleSupply int          `json:"minpresalesupply"`
	PresaleStartDate string       `json:"presalestartdate"`
	PresaleEndDate   string       `json:"presaleenddate"`
	DexListingDate   string       `json:"dexlistingdate"`
	Creator          string       `json:"creator"`
	CreatorUsername  string       `json:"creatorUsername"`
	Type             string       `json:"type"`
	Code             string       `json:"code"`
	Privacy          string       `json:"privacy"`
	DefaultModel     string       `json:"defaultmodel"`
	CreatedAt        time.Time    `json:"createdAt"`
}

type Distribution struct {
	Creator   int `json:"creator"`
	Curator   int `json:"curator"`
	Ecosystem int `json:"echosystem"`
	Promoter  int `json:"promoter"`
	Scout     int `json:"scout"`
}
