package chat

type SocketMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
