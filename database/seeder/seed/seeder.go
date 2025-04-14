package seed

import (
	"fmt"
	"github.com/gocanto/blog/database"
	"github.com/google/uuid"
	"time"
)

type Seeder struct {
	dbConn *database.Connection
}

func MakeSeeder(dbConnection *database.Connection) *Seeder {
	return &Seeder{
		dbConn: dbConnection,
	}
}

func (s *Seeder) SeedUsers() (database.User, database.User) {
	users := MakeUsersSeed(s.dbConn)

	UserA := users.Create(UsersAttrs{
		Username: "gocanto",
		Name:     "Gus",
		IsAdmin:  true,
	})

	UserB := users.Create(UsersAttrs{
		Username: "li",
		Name:     "liane",
		IsAdmin:  false,
	})

	return UserA, UserB
}

func (s *Seeder) SeedPosts(UserA, UserB database.User) []database.Post {
	posts := MakePostsSeed(s.dbConn)

	PostsA := posts.CreatePosts(PostsAttrs{
		AuthorID:    UserA.ID,
		Slug:        fmt.Sprintf("post-slug-%s", uuid.NewString()),
		Title:       fmt.Sprintf("Post %s title", uuid.NewString()),
		Excerpt:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
		Content:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris. Nullam tempor velit id mi luctus, a scelerisque libero accumsan. In hac habitasse platea dictumst. Cras ac nunc nec massa tristique fringilla.", uuid.NewString()),
		PublishedAt: time.Now(),
		Author:      UserA,
		Categories:  []database.Category{},
		Tags:        []database.Tag{},
		PostViews:   []database.PostView{},
		Comments:    []database.Comment{},
		Likes:       []database.Like{},
	}, 1)

	PostsB := posts.CreatePosts(PostsAttrs{
		AuthorID:    UserB.ID,
		Slug:        fmt.Sprintf("post-slug-%s", uuid.NewString()),
		Title:       fmt.Sprintf("Post %s title", uuid.NewString()),
		Excerpt:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
		Content:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris. Nullam tempor velit id mi luctus, a scelerisque libero accumsan. In hac habitasse platea dictumst. Cras ac nunc nec massa tristique fringilla.", uuid.NewString()),
		PublishedAt: time.Now(),
		Author:      UserB,
		Categories:  []database.Category{},
		Tags:        []database.Tag{},
		PostViews:   []database.PostView{},
		Comments:    []database.Comment{},
		Likes:       []database.Like{},
	}, 1)

	return append(PostsA, PostsB...)
}

func (s *Seeder) SeedCategories() {
	categories := MakeCategoriesSeed(s.dbConn)

	categories.Create(CategoriesAttrs{
		Slug:        fmt.Sprintf("category-slug-%s", uuid.NewString()),
		Description: fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
	})
}
