package seeds

import (
	"fmt"
	"github.com/gocanto/blog/database"
	"github.com/google/uuid"
	"math/rand"
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
	timex := time.Now()

	PostsA := posts.CreatePosts(PostsAttrs{
		AuthorID:    UserA.ID,
		Slug:        fmt.Sprintf("post-slug-%s", uuid.NewString()),
		Title:       fmt.Sprintf("Post %s title", uuid.NewString()),
		Excerpt:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
		Content:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris. Nullam tempor velit id mi luctus, a scelerisque libero accumsan. In hac habitasse platea dictumst. Cras ac nunc nec massa tristique fringilla.", uuid.NewString()),
		PublishedAt: &timex,
		Author:      UserA,
	}, 1)

	PostsB := posts.CreatePosts(PostsAttrs{
		AuthorID:    UserB.ID,
		Slug:        fmt.Sprintf("post-slug-%s", uuid.NewString()),
		Title:       fmt.Sprintf("Post %s title", uuid.NewString()),
		Excerpt:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
		Content:     fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris. Nullam tempor velit id mi luctus, a scelerisque libero accumsan. In hac habitasse platea dictumst. Cras ac nunc nec massa tristique fringilla.", uuid.NewString()),
		PublishedAt: &timex,
		Author:      UserB,
	}, 1)

	return append(PostsA, PostsB...)
}

func (s *Seeder) SeedCategories() []database.Category {
	categories := MakeCategoriesSeed(s.dbConn)

	return categories.Create(CategoriesAttrs{
		Slug:        fmt.Sprintf("category-slug-%s", uuid.NewString()),
		Description: fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
	})
}

func (s *Seeder) SeedTags() []database.Tag {
	seed := MakeTagsSeed(s.dbConn)

	return seed.Create()
}

func (s *Seeder) SeedComments(posts ...database.Post) {
	seed := MakeCommentsSeed(s.dbConn)

	timex := time.Now()
	var attrs []CommentsAttrs

	for index, post := range posts {
		attrs = append(attrs, CommentsAttrs{
			PostID:     post.ID,
			AuthorID:   post.AuthorID,
			ParentID:   nil,
			Content:    fmt.Sprintf("[%d] Nullam quis arcu in magna pulvinar tincidunt. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam hendrerit nulla ut cursus laoreet. Nullam elementum lorem vel facilisis laoreet. Cras ac turpis vel erat vehicula venenatis.", index),
			ApprovedAt: &timex,
		})
	}

	if _, err := seed.Create(attrs...); err != nil {
		panic(err)
	}
}

func (s *Seeder) SeedLikes(posts ...database.Post) {
	seed := MakeLikesSeed(s.dbConn)
	var attrs []LikesAttrs

	for _, post := range posts {
		attrs = append(attrs, LikesAttrs{
			PostID: post.ID,
			UserID: post.AuthorID,
		})
	}

	seed.Create(attrs...)
}

func (s *Seeder) SeedPostsCategories(categories []database.Category, posts []database.Post) {
	if len(categories) == 0 || len(posts) == 0 {
		return
	}

	seed := MakePostCategoriesSeed(s.dbConn)

	var post database.Post
	var category database.Category

	source := rand.NewSource(time.Now().UnixNano())
	salt := rand.New(source)

	cIndex := salt.Intn(len(categories))
	category = categories[cIndex]

	pIndex := salt.Intn(len(posts))
	post = posts[pIndex]

	seed.Create(category, post)

}

func (s *Seeder) SeedPostTags(tags []database.Tag, posts []database.Post) {
	if len(tags) == 0 || len(posts) == 0 {
		return
	}

	seed := MakePostTagsSeed(s.dbConn)

	var post database.Post
	var label database.Tag

	source := rand.NewSource(time.Now().UnixNano())
	salt := rand.New(source)

	tIndex := salt.Intn(len(tags))
	label = tags[tIndex]

	pIndex := salt.Intn(len(posts))
	post = posts[pIndex]

	seed.Create(label, post)

}

func (s *Seeder) SeedPostViews(posts []database.Post, users ...database.User) {
	if len(posts) == 0 || len(users) == 0 {
		return
	}

	seed := MakePostViewsSeed(s.dbConn)

	var attrs []PostViewsAttr

	for pIndex, post := range posts {
		for uIndex, user := range users {
			attrs = append(attrs, PostViewsAttr{
				Post:      post,
				User:      user,
				IPAddress: fmt.Sprintf("192.168.0.%d", pIndex+1),
				UserAgent: fmt.Sprintf("[post:%d][user:%d] - Mozilla/5.0 (Macintosh; ...) ...", pIndex+1, uIndex+1),
			})
		}
	}

	seed.Create(attrs)
}
