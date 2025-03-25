package agents

import (
	"net/http"

	agentsApp "github.com/mmosh-pit/mmosh_backend/pkg/agents/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetAgentsHandler(w http.ResponseWriter, r *http.Request) {
	agents := agentsApp.GetAgents()

	common.SendSuccessResponse(w, http.StatusOK, agents)
}
