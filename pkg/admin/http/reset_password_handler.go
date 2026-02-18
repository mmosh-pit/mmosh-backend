package http

import (
	"log"
	"net/http"

	adminApp "github.com/mmosh-pit/mmosh_backend/pkg/admin/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {

	userId := r.PathValue("userId")
	if userId == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid-parameter")
		return
	}

	err := adminApp.DeactivateUser(userId)

	if err != nil {
		log.Printf("[ADMIN/RESET PASSWORD] could not reset password: %v\n", err)

		common.SendErrorResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
