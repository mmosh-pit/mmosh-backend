package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
)

func SetOnboardingStepHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	var params authDomain.OnboardingStepParams

	err = json.Unmarshal(body, &params)
	if err != nil {
		log.Printf("error decoding payload on create guest user data: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"invalid payload"})
		return
	}

	auth.SetOnboardingStep(userId, params.Step)

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
