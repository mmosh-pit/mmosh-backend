package chat

import (
	"net/http"

	chatApp "github.com/mmosh-pit/mmosh_backend/pkg/chat/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetChatHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	response := chatApp.GetOrCreateChatForUser(userId)

	common.SendSuccessResponse(w, http.StatusOK, response)
}
