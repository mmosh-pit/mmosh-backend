package auth

import (
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func DeleteTelegramHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	auth.DeleteTelegramData(userId)

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
