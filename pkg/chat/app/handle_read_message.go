package chat

import (
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func HandleReadMessage(message chatDomain.SocketMessage, c *ClientData) {

	log.Printf("Received message: %v\n", message)

	switch message.Event {
	case "message":

		message, err := common.DecodeMapData[Message](message.Data)

		if err != nil {
			log.Printf("Got error trying to decode message: %v\n", err)
			break
		}

		message.UserId = c.UserId

		log.Println("sending message to pool")

		c.Client.Pool.SendMessage <- &message
		break
	}
}
