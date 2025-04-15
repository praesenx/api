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

    //fmt.Println("==========>", PostsA, PostsB)

    return append(PostsA, PostsB...)
}

func (s *Seeder) SeedCategories() {
    categories := MakeCategoriesSeed(s.dbConn)

    categories.Create(CategoriesAttrs{
        Slug:        fmt.Sprintf("category-slug-%s", uuid.NewString()),
        Description: fmt.Sprintf("[%s] Sed at risus vel nulla consequat fermentum. Donec et orci mauris", uuid.NewString()),
    })
}

func (s *Seeder) SeedTags() {
    seed := MakeTagsSeed(s.dbConn)
    seed.Create()
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
