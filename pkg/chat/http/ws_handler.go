package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	chatApp "github.com/mmosh-pit/mmosh_backend/pkg/chat/app"
)

var upgrader = websocket.Upgrader{}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := r.Header.Get("userId")

	token := r.URL.Query().Get("token")

	poolClient := &chatApp.PoolClient{
		Conn:  conn,
		Pool:  chatApp.WsPool,
		Token: token,
	}

	clientData := &chatApp.ClientData{
		Client: poolClient,
		UserId: user,
	}

	chatApp.WsPool.Connect <- clientData
	clientData.Read()
}
