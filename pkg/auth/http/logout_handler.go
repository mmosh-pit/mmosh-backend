package auth

import (
	"log"
	"net/http"
	"strings"

	authApp "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("userId")

	reqToken := r.Header.Get("Authorization")
	token := strings.Replace(reqToken, "Bearer ", "", 1)

	err := authApp.Logout(userId, token)
	if err != nil {
		log.Printf("error logout: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
