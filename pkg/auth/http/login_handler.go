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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	var data authDomain.LoginParams
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	response, err := authApp.Login(data)
	if err != nil {
		log.Printf("error login: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	log.Printf("Returning login: %v\n", *response.Token)

	common.SendSuccessResponse(w, http.StatusOK, response)
}
