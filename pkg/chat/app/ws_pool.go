package chat

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pool struct {
	Connect     chan *ClientData
	Leave       chan *ClientData
	AskAI       chan *AIMessage
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
	Pool       *Pool
	Token      string
}

type AIMessage struct {
	UserId       string
	Content      string
	ChatId       string
	AgentId      string
	Namespaces   []string
	SystemPrompt string
}

type Message struct {
	ChatId       string   `json:"chat_id"`
	AgentId      string   `json:"agent_id"`
	UserId       string   `json:"user_id"`
	Content      string   `json:"content"`
	Namespaces   []string `json:"namespaces"`
	SystemPrompt string   `json:"system_prompt"`
}

var WsPool *Pool

func CreatePool() {
	WsPool = &Pool{
		Connect:     make(chan *ClientData),
		Leave:       make(chan *ClientData),
		Clients:     make(map[string]*PoolClient),
		AskAI:       make(chan *AIMessage),
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

			chatId, err := primitive.ObjectIDFromHex(message.ChatId)

			agentId, err := primitive.ObjectIDFromHex(message.AgentId)

			resultingId := &agentId

			if err != nil {
				resultingId = nil
			}

			messageData := chat.Message{
				ID:        &id,
				Content:   message.Content,
				Sender:    &userId,
				CreatedAt: time.Now(),
				Type:      "user",
				AgentId:   resultingId,
				ChatId:    &chatId,
			}

			data := map[string]any{
				"event": "userMessage",
				"data":  messageData,
			}

			go GenerateAIResponse(client, &messageData)

			client.sendResponse(data)
			break

		case message := <-WsPool.SendMessage:
			client, ok := WsPool.Clients[message.UserId]
			if !ok {
				break
			}

			log.Println("Gonna send...")

			id := primitive.NewObjectID()

			userId, _ := primitive.ObjectIDFromHex(message.UserId)

			agentId, err := primitive.ObjectIDFromHex(message.AgentId)

			resultingId := &agentId

			if err != nil {
				resultingId = nil
			}

			chatId, err := primitive.ObjectIDFromHex(message.ChatId)

			if err != nil {
				log.Printf("Invalid chat object ID: %v, %v\n", message.ChatId, err)
				break
			}

			messageData := chat.Message{
				ID:           &id,
				Content:      message.Content,
				Sender:       &userId,
				CreatedAt:    time.Now(),
				Type:         "user",
				Namespaces:   message.Namespaces,
				SystemPrompt: message.SystemPrompt,
				AgentId:      resultingId,
				ChatId:       &chatId,
			}

			data := map[string]any{
				"event": "userMessage",
				"data":  messageData,
			}
			log.Println("6")

			go GenerateAIResponse(client, &messageData)
			log.Println("7")

			client.sendResponse(data)
			break
		}

	}
}

func (p *PoolClient) sendResponse(message any) {
	p.WriteMutex.Lock()
	defer p.WriteMutex.Unlock()
	err := p.Conn.WriteJSON(message)

	if err != nil {
		log.Printf("Error while trying to send JSON: %v\n", err)
	}
}
