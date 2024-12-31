package chat

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
)

func (c *ClientData) Read() {
	defer func() {
		WsPool.Leave <- c
		err := c.Client.Conn.Close()

		if err != nil {
			log.Printf("Error closing connection: %s", err)
		}
	}()

	KeepWsAlive(c.Client, time.Second*10)

	for {
		_, data, err := c.Client.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error: %s", err)
			return
		}

		log.Printf("We got data: %v\n", string(data))

		var decoded chat.SocketMessage

		err = json.Unmarshal(data, &decoded)

		if err != nil {
			log.Printf("Error decoding incoming socket message: %v\n %v\n", string(data), err)

			return
		}

		HandleReadMessage(decoded, c)
	}
}

func KeepWsAlive(c *PoolClient, timeout time.Duration) {
	time.Sleep(timeout / 2)
	lastResponse := time.Now()
	c.Conn.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		for {
			err := c.Conn.WriteMessage(websocket.PingMessage, []byte("ka"))
			if err != nil {
				return
			}
			time.Sleep(timeout / 2)
			if time.Since(lastResponse) > timeout {
				log.Println("Closing by timeout...")
				c.Conn.Close()
				return
			}
		}
	}()
}
