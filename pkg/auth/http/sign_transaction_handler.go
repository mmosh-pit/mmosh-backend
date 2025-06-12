package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func SignTransactionHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	var data auth.SignTransactionRequest

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload on signup: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	res, err := auth.SignTransaction(userId, data.Message)

	if err != nil {
		log.Printf("Got error: %v\n", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, []string{"something-went-wrong"})
		return
	}

	common.SendSuccessResponse(w, http.StatusCreated, res)
}
