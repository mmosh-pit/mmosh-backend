package chat

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pool struct {
	Connect     chan *ClientData
	Leave       chan *ClientData
	AskAI       chan *Message
	SendMessage chan *Message
	Clients     map[string]*PoolClient
}

type ClientData struct {
	Client *PoolClient
	UserId string
}

type PoolClient struct {
	Conn       *websocket.Conn
	WriteMutex sync.Mutex
}

type Message struct {
	UserId  string
	Content string
}

var WsPool *Pool

func CreatePool() {
	WsPool = &Pool{
		Connect:     make(chan *ClientData),
		Leave:       make(chan *ClientData),
		Clients:     make(map[string]*PoolClient),
		AskAI:       make(chan *Message),
		SendMessage: make(chan *Message),
	}

	Start()
}

func Start() {
	for {
		select {
		case client := <-WsPool.Connect:
			WsPool.Clients[client.UserId] = client.Client
			log.Println("Size of Connection Pool: ", len(WsPool.Clients))

			// client.Client.SendSimpleResponse("Connected")
			client.Client.Conn.WriteJSON("connected")
			break

		case client := <-WsPool.Leave:
			delete(WsPool.Clients, client.UserId)
			break

		case message := <-WsPool.AskAI:
			client, ok := WsPool.Clients[message.UserId]
			if !ok {
				break
			}

			id := primitive.NewObjectID()

			userId, _ := primitive.ObjectIDFromHex(message.UserId)

			data := chat.Message{
				ID:      &id,
				Content: message.Content,
				Sender:  &userId,
				Type:    "user",
			}

			go GenerateAIResponse(client, &data)

			client.sendResponse(data)
			break
		}

	}
}

func (p *PoolClient) sendResponse(message interface{}) {
	p.WriteMutex.Lock()
	defer p.WriteMutex.Unlock()
	p.Conn.WriteJSON(message)
}
