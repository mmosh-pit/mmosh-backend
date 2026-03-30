package members

import (
	"net/http"
	"strconv"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	members "github.com/mmosh-pit/mmosh_backend/pkg/members/db"
)

func GetMembersHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, "")
		return
	}

	search := r.URL.Query().Get("search")

	users := members.GetMembers(int64(page), search, userId)

	common.SendSuccessResponse(w, http.StatusOK, users)
}
