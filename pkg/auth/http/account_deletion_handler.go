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

func AccountDeletionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	var data authDomain.AccountDeletionRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload on signup: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	err = auth.AccountDeletion(&data)

	if err != nil {
		log.Printf("error requesting account deletion: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
