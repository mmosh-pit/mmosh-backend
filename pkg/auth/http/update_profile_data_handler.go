package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func UpdateProfileDataHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	var data auth.UpdateProfileDataParams
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload on signup: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	err = auth.UpdateProfileData(data, userId)

	if err != nil {
		log.Printf("error updating profile data: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
