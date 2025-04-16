package seeds

import (
	"github.com/gocanto/blog/database"
	"github.com/google/uuid"
)

type LikesSeed struct {
	db *database.Connection
}

type LikesAttrs struct {
	UUID   string `gorm:"type:uuid;unique;not null"`
	PostID uint64 `gorm:"not null;index;uniqueIndex:idx_likes_post_user"`
	UserID uint64 `gorm:"not null;index;uniqueIndex:idx_likes_post_user"`
}

func MakeLikesSeed(db *database.Connection) *LikesSeed {
	return &LikesSeed{
		db: db,
	}
}

func (s LikesSeed) Create(attrs ...LikesAttrs) []database.Like {
	var likes []database.Like

	for _, attr := range attrs {
		likes = append(likes, database.Like{
			UUID:   uuid.NewString(),
			PostID: attr.PostID,
			UserID: attr.UserID,
		})
	}

	s.db.Sql().Create(&likes)

	return likes
}
