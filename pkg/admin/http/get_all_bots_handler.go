package http

import (
	"net/http"
	"strconv"

	adminApp "github.com/mmosh-pit/mmosh_backend/pkg/admin/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetAllBotsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	search := query.Get("search")

	page, err := strconv.Atoi(query.Get("page"))

	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	bots := adminApp.GetAllBots(page, search)

	common.SendSuccessResponse(w, http.StatusOK, bots)
}
