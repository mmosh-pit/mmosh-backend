package agents

import (
	"errors"
	"net/http"

	agentsApp "github.com/mmosh-pit/mmosh_backend/pkg/agents/app"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/agents/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetActiveAgentsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	agents, err := agentsApp.GetActiveAgents(userId)

	if err != nil {
		switch {
		case errors.Is(err, agentsDomain.ErrUserNotFound):
			common.SendErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
			return

		default:
			common.SendErrorResponse(w, http.StatusInternalServerError, []string{"something-went-wrong"})
			return
		}
	}

	common.SendSuccessResponse(w, http.StatusOK, agents)
}
