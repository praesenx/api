package seeds

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/oullin/database"
	"github.com/oullin/pkg/gorm"
	"time"
)

type PostsSeed struct {
	db *database.Connection
}

type PostsAttrs struct {
	AuthorID    uint64
	Slug        string
	Title       string
	Excerpt     string
	Content     string
	PublishedAt *time.Time
	Author      database.User
	Categories  []database.Category
	Tags        []database.Tag
	PostViews   []database.PostView
	Comments    []database.Comment
	Likes       []database.Like
}

func MakePostsSeed(db *database.Connection) *PostsSeed {
	return &PostsSeed{
		db: db,
	}
}

func (s PostsSeed) CreatePosts(attrs PostsAttrs, number int) ([]database.Post, error) {
	var posts []database.Post

	for i := 1; i <= number; i++ {
		post := database.Post{
			UUID:          uuid.NewString(),
			AuthorID:      attrs.AuthorID,
			Slug:          fmt.Sprintf("%s-post-%s-%d", attrs.Author.Username, attrs.Slug, i),
			Title:         fmt.Sprintf("Post: [%d] by %s", i, attrs.Author.Username),
			Excerpt:       "This is an excerpt.",
			Content:       "This is the full content of the post.",
			CoverImageURL: "",
			PublishedAt:   attrs.PublishedAt,
			Categories:    []database.Category{},
			Tags:          []database.Tag{},
			PostViews:     []database.PostView{},
			Comments:      []database.Comment{},
			Likes:         []database.Like{},
		}

		posts = append(posts, post)
	}

	result := s.db.Sql().Create(&posts)

	if gorm.HasDbIssues(result.Error) {
		return nil, fmt.Errorf("issue creating posts: %s", result.Error)
	}

	return posts, nil
}
