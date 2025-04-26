package posts

import (
	"log"
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	postsApp "github.com/mmosh-pit/mmosh_backend/pkg/posts/app"
)

func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {

	posts, err := postsApp.GetAllPosts()
	if err != nil {
		log.Printf("ERROR: Failed to retrieve all posts: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, []string{"Internal Server Error"})
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, posts)
}
