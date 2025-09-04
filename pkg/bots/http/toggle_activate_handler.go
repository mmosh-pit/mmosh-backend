package bots

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	agentsApp "github.com/mmosh-pit/mmosh_backend/pkg/bots/app"
	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func ToggleActivateHandler(w http.ResponseWriter, r *http.Request) {
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

	var data agentsDomain.ToggleActivateAgentData

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	err = agentsApp.ActivateDeactivateAgent(userId, data.AgentId)

	if err != nil {
		log.Printf("[ACTIVATING AGENTS] error: %v\n", err)
		switch {
		case errors.Is(err, agentsDomain.ErrAgentNotExists), errors.Is(err, agentsDomain.ErrUserNotSubscribed), errors.Is(err, agentsDomain.ErrUserNotFound):

			common.SendErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		default:
			common.SendErrorResponse(w, http.StatusInternalServerError, "something-went-wrong")
			return
		}

	}

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
