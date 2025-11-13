package chat

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func SendBotMessageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	var data chat.SendBotMessageData

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload on send bot message handler: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	err = chat.SendBotMessage(data)

}
