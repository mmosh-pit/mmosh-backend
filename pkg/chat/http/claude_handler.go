package chat

import (
	"encoding/json"
	"log"
	"net/http"

	chatApp "github.com/mmosh-pit/mmosh_backend/pkg/chat/app"
	domain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func ClaudeHandler(w http.ResponseWriter, r *http.Request) {
	// ── Parse request ─────────────────────────────────────────────────────────
	var req domain.ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid request")
		return
	}

	if req.Message == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "message is required")
		return
	}

	key := config.GetAnthropicKey()

	claudeService := chatApp.NewService(key)

	// ── Call Claude ───────────────────────────────────────────────────────────
	resp, err := claudeService.Chat(r.Context(), &req)
	if err != nil {
		log.Printf("Got error here in chat request for claude: %v\n", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "failed to process chat request")
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, resp)
}
