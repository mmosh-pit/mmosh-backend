package auth

import (
	"log"
	"net/http"

	authApp "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func IsAuthHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	response, err := authApp.RetrieveUserById(userId)
	if err != nil {
		log.Printf("Error checking if is authenticated: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, response)
}
