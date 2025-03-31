package agents

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	agentsApp "github.com/mmosh-pit/mmosh_backend/pkg/agents/app"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/agents/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func CreateAgentHandler(w http.ResponseWriter, r *http.Request) {
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

	var data agentsDomain.CreateAgentData

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	err = agentsApp.CreateAgent(data, userId)

	if err != nil {
		switch {
		case errors.Is(err, agentsDomain.ErrAgentNotExists), errors.Is(err, agentsDomain.ErrUserNotSubscribed), errors.Is(err, agentsDomain.ErrUserNotFound):

			common.SendErrorResponse(w, http.StatusBadRequest, []string{err.Error()})
			return
		default:
			common.SendErrorResponse(w, http.StatusInternalServerError, []string{"something-went-wrong"})
			return
		}

	}

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
