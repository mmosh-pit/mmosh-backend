package http

import (
	"net/http"
	"strconv"

	adminApp "github.com/mmosh-pit/mmosh_backend/pkg/admin/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	page, err := strconv.Atoi(query.Get("page"))

	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid-request")
		return
	}

	search := query.Get("search")

	users := adminApp.GetAllUsers(int64(page), search)

	common.SendSuccessResponse(w, http.StatusOK, users)
}
