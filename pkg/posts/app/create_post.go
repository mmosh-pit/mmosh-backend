package posts

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	commonApp "github.com/mmosh-pit/mmosh_backend/pkg/common/app"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
)

const SYSTEM_PROMPT = `
  Your job is to post content to the feed. The user will send you the headline and the body of the post. The user may also tell you if they want you to post it as-is, or if they want you to edit it or rewrite it.

  If it’s not clear, ask clarifying questions until you’re confident that you know what what to post, where to post it and if you should edit it, rewrite it completely or post as-is.

  Please when generating the post, divide it into sections by using their name, separated by {section_name} using "{{}}" brackets, instead of astericks.
  Possible sections: Headline, Body, Tags
`

// use the style for the bot that you’re posting this for.

type CreatePostParams struct {
	Prompt string `json:"prompt"`
	UserId string
}

func CreatePost(params *CreatePostParams) (*postsDomain.Post, error) {

	user, _ := auth.GetUserById(params.UserId)

	aiResponse := commonApp.FetchAIResponse(user.Name, params.Prompt, SYSTEM_PROMPT, []string{"PUBLIC", "MMOSH"})

	log.Printf("Got AI Response here: %v\n", aiResponse)

	post := postsDomain.Post{}
	//
	// post.Slug = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(post.Slug), " ", "-"))
	//
	// // if post.Slug == "" && post.Header != "" {
	// // 	post.Slug = slug.Make(post.Header)
	// // }
	//
	// if post.Slug == "" {
	// 	return nil, errors.New("slug cannot be empty or contain only spaces after formatting")
	// }
	//
	// post.CreatedAt = time.Now()
	// post.UpdatedAt = time.Now()
	//
	// err := postsDb.CreatePost(&post)
	// if err != nil {
	// 	return nil, fmt.Errorf("service error creating post: %w", err)
	// }
	return &post, nil
}
