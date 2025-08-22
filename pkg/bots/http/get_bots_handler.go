package bots

import (
	"net/http"
	"strconv"

	bots "github.com/mmosh-pit/mmosh_backend/pkg/bots/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetBotsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	query := r.URL.Query()

	search := query.Get("search")

	page, err := strconv.Atoi(query.Get("page"))

	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	bots := bots.GetBots(userId, search, page)

	common.SendSuccessResponse(w, http.StatusOK, bots)
}
