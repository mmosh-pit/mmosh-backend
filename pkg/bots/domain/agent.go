package bots

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bot struct {
	Id               *primitive.ObjectID `bson:"_id" json:"id"`
	Name             string              `bson:"name" json:"name"`
	Desc             string              `bson:"desc" json:"desc"`
	Image            string              `bson:"image" json:"image"`
	Symbol           string              `bson:"symbol" json:"symbol"`
	Key              string              `bson:"key" json:"key"`
	Price            int                 `bson:"price" json:"price"`
	PreSaleStartDate string              `bson:"presalestartdate" json:"presalestartdate"`
	SystemPrompt     string              `bson:"system_prompt" json:"system_prompt"`
	CreatorUsername  string              `bson:"creatorUsername" json:"creatorUsername"`
	Type             string              `bson:"type" json:"type"`
}

type ToggleActivateAgentData struct {
	AgentId string `json:"agentId"`
}

type CreateAgentData struct {
	Id               *primitive.ObjectID `bson:"_id" json:"id"`
	Name             string              `bson:"name" json:"name"`
	Symbol           string              `bson:"symbol" json:"symbol"`
	Desc             string              `bson:"desc" json:"desc"`
	Image            string              `bson:"image" json:"image"`
	InviteImage      string              `bson:"inviteimage" json:"inviteimage"`
	Key              string              `bson:"key" json:"key"`
	Lut              string              `bson:"lut" json:"lut"`
	Seniority        int                 `bson:"seniority" json:"seniority"`
	Price            float64             `bson:"price" json:"price"` // Assuming fields.passPrice can be a float
	Distribution     Distribution        `bson:"distribution" json:"distribution"`
	InvitationPrice  float64             `bson:"invitationprice" json:"invitationprice"` // Assuming fields.invitationPrice can be a float
	Discount         float64             `bson:"discount" json:"discount"`               // Assuming fields.discount can be a float
	Telegram         string              `bson:"telegram" json:"telegram"`
	Twitter          string              `bson:"twitter" json:"twitter"`
	Website          string              `bson:"website" json:"website"`
	PresaleSupply    int                 `bson:"presalesupply" json:"presalesupply"`
	MinPresaleSupply int                 `bson:"minpresalesupply" json:"minpresalesupply"`
	PresaleStartDate string              `bson:"presalestartdate" json:"presalestartdate"`
	PresaleEndDate   string              `bson:"presaleenddate" json:"presaleenddate"`
	DexListingDate   string              `bson:"dexlistingdate" json:"dexlistingdate"`
	Creator          string              `bson:"creator" json:"creator"`
	CreatorUsername  string              `bson:"creatorUsername" json:"creatorUsername"`
	Type             string              `bson:"type" json:"type"`
	Code             string              `bson:"code" json:"code"`
	Privacy          string              `bson:"privacy" json:"privacy"`
}

type Distribution struct {
	Creator   int `bson:"creator" json:"creator"`
	Curator   int `bson:"curator" json:"curator"`
	Ecosystem int `bson:"ecosystem" json:"echosystem"`
	Promoter  int `bson:"promoter" json:"promoter"`
	Scout     int `bson:"scout" json:"scout"`
}
