package members

import (
	"net/http"
	"strconv"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	members "github.com/mmosh-pit/mmosh_backend/pkg/members/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, "")
		return
	}

	users := members.GetMembers(int64(page), search, userIdBson)

	common.SendSuccessResponse(w, http.StatusOK, users)
}
