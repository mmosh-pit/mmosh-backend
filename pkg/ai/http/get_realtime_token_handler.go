package ai

import (
	"net/http"

	ai "github.com/mmosh-pit/mmosh_backend/pkg/ai/app"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func GetRealtimeTokenHandler(w http.ResponseWriter, r *http.Request) {

	result, err := ai.GetRealtimeToken()

	if err != nil {

		if err.Error() == "unauthorized" {
			common.SendErrorResponse(w, http.StatusForbidden, "unauthorized")
			return
		}

		common.SendErrorResponse(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, result)
}
