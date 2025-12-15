package http

import (
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	configApp "github.com/mmosh-pit/mmosh_backend/pkg/config/app"
)

func GetAppThemesHandler(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	result := configApp.GetAppThemes()

	common.SendSuccessResponse(w, http.StatusOK, result)
}
