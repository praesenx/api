package seeds

import (
    "github.com/gocanto/blog/database"
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

func (s PostViewsSeed) Create(attrs []PostViewsAttr) {
    for _, attr := range attrs {
        s.db.Sql().Create(&database.PostView{
            PostID:    attr.Post.ID,
            UserID:    &attr.User.ID,
            IPAddress: attr.IPAddress,
            UserAgent: attr.UserAgent,
        })
    }
}
