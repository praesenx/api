package seed

import (
	"github.com/gocanto/blog/database"
)

type PostTagsSeed struct {
	db *database.Connection
}

func MakePostTagsSeed(db *database.Connection) *PostTagsSeed {
	return &PostTagsSeed{
		db: db,
	}
}

func (s PostTagsSeed) Create(tag database.Tag, post database.Post) {
	s.db.Sql().Create(&database.PostTag{
		PostID: post.ID,
		TagID:  tag.ID,
	})
}
