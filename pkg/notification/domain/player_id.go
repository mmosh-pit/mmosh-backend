package notificationDomain

type InsertPlayIdRequestParams struct {
	PlayerId string `json:"playerId" bson:"playerId"`
	Wallet   string `json:"wallet" bson:"wallet"`
}
