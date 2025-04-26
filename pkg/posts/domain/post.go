package posts

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	Header string `bson:"header" json:"header"`

	SubHeader string `bson:"subHeader,omitempty" json:"subHeader,omitempty"`

	Tags []string `bson:"tags,omitempty" json:"tags,omitempty"`

	Authors []string `bson:"authors" json:"authors"`

	Body string `bson:"body" json:"body"`

	Slug string `bson:"slug" json:"slug"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt,omitempty"`
}

func (p *Post) Validate() bool {
	headerValid := strings.TrimSpace(p.Header) != ""
	bodyValid := strings.TrimSpace(p.Body) != ""
	slugValid := strings.TrimSpace(p.Slug) != ""
	authorsValid := len(p.Authors) > 0

	if authorsValid {
		for _, author := range p.Authors {
			if strings.TrimSpace(author) == "" {
				authorsValid = false
				break
			}
		}
	}

	return headerValid && bodyValid && slugValid && authorsValid
}
