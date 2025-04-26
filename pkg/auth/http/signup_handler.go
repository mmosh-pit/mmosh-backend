package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	authApp "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	var data authDomain.SignUpParams
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload on signup: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	response, err := authApp.SignUp(&data)
	if err != nil {

		switch err {
		case authDomain.ErrWalletAlreadyExists:
			common.SendErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
			return
		default:
			log.Printf("error sign up: %v", err)
			common.SendErrorResponse(w, http.StatusInternalServerError, []string{err.Error()})
			return
		}
	}

	common.SendSuccessResponse(w, http.StatusOK, response)
}
