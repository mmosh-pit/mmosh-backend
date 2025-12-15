package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	configApp "github.com/mmosh-pit/mmosh_backend/pkg/config/app"
	configDomain "github.com/mmosh-pit/mmosh_backend/pkg/config/domain"
)

func CreateAppThemeHandler(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("[CREATE APP THEME] error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	var data configDomain.AppTheme

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("[CREATE APP THEME] error decoding payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	err = configApp.CreateAppTheme(data, userId)

	if err != nil {
		log.Printf("[CREATE APP THEME] could not create theme: %v\n", err)

		if err == configDomain.ErrUserNotAuthorized {
			common.SendErrorResponse(w, http.StatusForbidden, err.Error())
			return
		}

		common.SendErrorResponse(w, http.StatusInternalServerError, "something-went-wrong")
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
