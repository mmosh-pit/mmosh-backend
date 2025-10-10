package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	authApp "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func ForgotPasswordVerificationHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	var data authApp.RequestCodeParams

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload on request code: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	err = authApp.RequestCode(data.Email)
	if err != nil {
		log.Printf("error request code: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, "")
}
