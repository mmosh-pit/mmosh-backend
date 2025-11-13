package chat

import (
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func HandleReadMessage(message chatDomain.SocketMessage, c *ClientData) {

	switch message.Event {
	case "message":

		message, err := common.DecodeMapData[Message](message.Data)

		if err != nil {
			log.Printf("Got error trying to decode message: %v\n", err)
			break
		}

		message.UserId = c.UserId

		c.Client.Pool.SendMessage <- &message
		break
	}
}
