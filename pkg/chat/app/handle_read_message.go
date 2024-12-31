package chat

import (
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
)

func HandleReadMessage(message chatDomain.SocketMessage, c *ClientData) {
	switch message.Event {
	case "message":
		WsPool.AskAI <- &Message{
			Content: message.Data.(string),
			UserId:  c.UserId,
		}
		break
	}
}
