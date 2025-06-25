package bots

import (
	"net/http"

	bots "github.com/mmosh-pit/mmosh_backend/pkg/bots/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetMyBotsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	bots := bots.GetMyBots(userId)

	common.SendSuccessResponse(w, http.StatusOK, bots)
}
