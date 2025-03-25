package chat

import (
	"errors"
	"net/http"

	chatApp "github.com/mmosh-pit/mmosh_backend/pkg/chat/app"
	commonDomain "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetActiveChatsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	response, err := chatApp.GetActiveChats(userId)

	if err != nil {
		switch {
		case errors.Is(err, commonDomain.UserNotExistsErr):
			common.SendErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
			return
		default:
			common.SendErrorResponse(w, http.StatusInternalServerError, []string{"something-went-wrong"})
			return
		}
	}

	common.SendSuccessResponse(w, http.StatusOK, response)
}
