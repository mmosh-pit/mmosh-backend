package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func CreateGuestUserDataHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error receiving payload: %v\n", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	var params authDomain.GuestUserData

	err = json.Unmarshal(body, &params)
	if err != nil {
		log.Printf("error decoding payload on create guest user data: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	err = auth.CreateGuestUserData(params, userId)

	if err != nil {
		switch err {
		default:
			common.SendErrorResponse(w, http.StatusInternalServerError, []string{err.Error()})
			return
		}
	}

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
