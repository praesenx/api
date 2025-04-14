package seed

import (
	"fmt"
	"github.com/gocanto/blog/database"
	"github.com/google/uuid"
	"time"
)

type PostSeed struct {
	db *database.Connection
}

type PostsAttrs struct {
	AuthorID    uint64
	Slug        string
	Title       string
	Excerpt     string
	Content     string
	PublishedAt time.Time
	Author      database.User
	Categories  []database.Category
	Tags        []database.Tag
	PostViews   []database.PostView
	Comments    []database.Comment
	Likes       []database.Like
}

func MakePostsSeed(db *database.Connection) *PostSeed {
	return &PostSeed{
		db: db,
	}
}

func (s PostSeed) CreatePosts(attrs PostsAttrs, number int) []database.Post {
	var posts []database.Post

	for i := 1; i <= number; i++ {
		post := database.Post{
			UUID:          uuid.NewString(),
			AuthorID:      attrs.AuthorID,
			Slug:          fmt.Sprintf("%s-post-%s-%d", attrs.Author.Username, attrs.Slug, i),
			Title:         fmt.Sprintf("Post %d by %s", i, attrs.Author.Username),
			Excerpt:       "This is an excerpt.",
			Content:       "This is the full content of the post.",
			CoverImageURL: "",
			PublishedAt:   nil, //time.Now(),
			Categories:    []database.Category{},
			Tags:          []database.Tag{},
			PostViews:     []database.PostView{},
			Comments:      []database.Comment{},
			Likes:         []database.Like{},
		}

		posts = append(posts, post)
	}

	s.db.Sql().Create(&posts)

	return posts
}
