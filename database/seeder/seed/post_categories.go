package seed

import (
    "github.com/gocanto/blog/database"
)

type PostCategoriesSeed struct {
    db *database.Connection
}

func MakePostCategoriesSeed(db *database.Connection) *PostCategoriesSeed {
    return &PostCategoriesSeed{
        db: db,
    }
}

func (s PostCategoriesSeed) Create(category database.Category, post database.Post) {
    s.db.Sql().Create(&database.PostCategory{
        CategoryID: category.ID,
        PostID:     post.ID,
    })
}
