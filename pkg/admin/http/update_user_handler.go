package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	adminApp "github.com/mmosh-pit/mmosh_backend/pkg/admin/app"
	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	if userId == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid-parameter")
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	var data adminDomain.UpdateUserPayload

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	adminApp.UpdateUser(data)

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
