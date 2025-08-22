package bots

import (
	"errors"
	"net/http"

	agentsApp "github.com/mmosh-pit/mmosh_backend/pkg/bots/app"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetActiveAgentsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	query := r.URL.Query()

	search := query.Get("search")

	agents, err := agentsApp.GetActiveAgents(userId, search)

	if err != nil {
		switch {
		case errors.Is(err, agentsDomain.ErrUserNotFound):
			common.SendErrorResponse(w, http.StatusBadRequest, err.Error())
			return

		default:
			common.SendErrorResponse(w, http.StatusInternalServerError, "something-went-wrong")
			return
		}
	}

	common.SendSuccessResponse(w, http.StatusOK, agents)
}
