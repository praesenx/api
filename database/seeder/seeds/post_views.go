package seeds

import (
	"fmt"
	"github.com/oullin/api/database"
	"github.com/oullin/api/pkg/gorm"
)

type PostViewsSeed struct {
	db *database.Connection
}

type PostViewsAttr struct {
	Post      database.Post
	User      database.User
	IPAddress string
	UserAgent string
}

func MakePostViewsSeed(db *database.Connection) *PostViewsSeed {
	return &PostViewsSeed{
		db: db,
	}
}

func (s PostViewsSeed) Create(attrs []PostViewsAttr) error {
	for _, attr := range attrs {
		result := s.db.Sql().Create(&database.PostView{
			PostID:    attr.Post.ID,
			UserID:    &attr.User.ID,
			IPAddress: attr.IPAddress,
			UserAgent: attr.UserAgent,
		})

		if gorm.HasDbIssues(result.Error) {
			return fmt.Errorf("issue creating post views for post: %s", result.Error)
		}
	}

	return nil
}
