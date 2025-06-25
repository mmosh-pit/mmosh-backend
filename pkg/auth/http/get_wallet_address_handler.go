package auth

import (
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetWalletAddressHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	wallet := auth.GetWalletAddres(userId)

	common.SendSuccessResponse(w, http.StatusOK, wallet)
}
