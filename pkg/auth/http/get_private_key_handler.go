package auth

import (
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetPrivateKeyHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	response, err := auth.GetUserPrivateKey(userId)

	if err != nil {
		common.SendErrorResponse(w, http.StatusInternalServerError, nil)
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, response)
}
