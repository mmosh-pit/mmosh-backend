package http

import (
	"log"
	"net/http"

	adminApp "github.com/mmosh-pit/mmosh_backend/pkg/admin/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func ActivateUserHandler(w http.ResponseWriter, r *http.Request) {

	userId := r.PathValue("userId")
	if userId == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid-parameter")
		return
	}

	err := adminApp.ActivateUser(userId)

	if err != nil {
		log.Printf("[ADMIN/ACTIVATE USER] could not activate user: %v\n", err)

		common.SendErrorResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
