package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	commonDomain "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func AddReferredToUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	var params authDomain.AddReferrerParams

	err = json.Unmarshal(body, &params)
	if err != nil {
		log.Printf("error decoding payload on adding referrer to user: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	err = auth.AddReferrerUser(params, userId)

	if err != nil {
		switch err {
		case commonDomain.UserNotExistsErr:
			common.SendErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
		default:
			common.SendErrorResponse(w, http.StatusInternalServerError, []string{err.Error()})
		}
	}

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
