package http

import (
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	common.SendSuccessResponse(w, http.StatusOK, nil)
}
