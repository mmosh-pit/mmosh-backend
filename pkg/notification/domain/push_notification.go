package notificationDomain

type PushNotificationRequestParams struct {
	Title   string `json:"title" bson:"title"`
	Message string `json:"message" bson:"message"`
	Wallet  string `json:"wallet" bson:"wallet"`
}

type OneSignalRequest struct {
	AppID            string            `json:"app_id"`
	IncludePlayerIDs []string          `json:"include_player_ids"`
	Headings         map[string]string `json:"headings"`
	Contents         map[string]string `json:"contents"`
}

type OneSignalResponse struct {
	ID         string      `json:"id"`
	Recipients int         `json:"recipients"`
	Errors     interface{} `json:"errors,omitempty"`
}
