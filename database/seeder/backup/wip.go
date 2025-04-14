package backup

import (
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"time"

	"github.com/gocanto/blog/database"
	//"github.com/gocanto/blog/env"
	"github.com/google/uuid"
)

//func run(env map[string]string) {
//	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
//
//	// Load environment (replace with actual loading if needed)
//	e := env.NewEnvironment()
//	conn, err := database.MakeConnection(e)
//	if err != nil {
//		log.Fatalf("failed to connect to DB: %v", err)
//	}
//	db := conn.Sql()
//
//	// Migrate the schema
//	db.AutoMigrate(&database.User{}, &database.Post{}, &database.Category{}, &database.Tag{},
//		&database.PostView{}, &database.Comment{}, &database.Like{})
//
//	// Create users
//	admin := createUser(db, "Admin", true)
//	regular := createUser(db, "User", false)
//
//	// Create posts
//	adminPosts := createPosts(db, admin, 3)
//	userPosts := createPosts(db, regular, 3)
//	allPosts := append(adminPosts, userPosts...)
//
//	// Create categories and link randomly
//	categories := createCategories(db)
//	linkCategoriesToPosts(db, allPosts, categories, rng)
//
//	// Create tags and link randomly
//	tags := createTags(db)
//	linkTagsToPosts(db, allPosts, tags)
//
//	// Create post views, comments, and likes
//	for _, post := range allPosts {
//		createViews(db, post, 10)
//		createComments(db, post, 5, []database.User{admin, regular})
//		createLikes(db, post, []database.User{admin, regular})
//	}
//
//	fmt.Println("✅ Fake data seeding complete.")
//}

//func createUser(db *gorm.DB, name string, isAdmin bool) database.User {
//	user := database.User{
//		UUID:         uuid.NewString(),
//		FirstName:    name,
//		LastName:     "Tester",
//		Username:     fmt.Sprintf("%sUser", name),
//		DisplayName:  fmt.Sprintf("%s User", name),
//		Email:        fmt.Sprintf("%s@test.com", name),
//		PasswordHash: "hashedpass",
//		PublicToken:  uuid.NewString(),
//		IsAdmin:      isAdmin,
//		VerifiedAt:   time.Now(),
//	}
//	db.Create(&user)
//	return user
//}

//func createPosts(db *gorm.DB, author database.User, count int) []database.Post {
//	posts := []database.Post{}
//	for i := 1; i <= count; i++ {
//		post := database.Post{
//			UUID:     uuid.NewString(),
//			AuthorID: author.ID,
//			Slug:     fmt.Sprintf("%s-post-%d", author.Username, i),
//			Title:    fmt.Sprintf("Post %d by %s", i, author.Username),
//			Excerpt:  "This is an excerpt.",
//			Content:  "This is the full content of the post.",
//		}
//		db.Create(&post)
//		posts = append(posts, post)
//	}
//	return posts
//}

func linkCategoriesToPosts(db *gorm.DB, posts []database.Post, cats []database.Category, rng *rand.Rand) {
	for _, post := range posts {
		n := rng.Intn(1)
		db.Model(&post).Association("Categories").Replace(randomSubset(cats, n, rng))
	}
}

func createTags(db *gorm.DB) []database.Tag {
	names := []string{"Tech", "AI", "Leadership", "Ethics", "Automation", "Teamwork", "Agile", "OpenAI", "Scaling", "Future"}
	var tags []database.Tag
	for _, name := range names {
		t := database.Tag{
			UUID: uuid.NewString(),
			Name: name,
			Slug: fmt.Sprintf("%s-tag", name),
		}
		db.Create(&t)
		tags = append(tags, t)
	}
	return tags
}

func linkTagsToPosts(db *gorm.DB, posts []database.Post, tags []database.Tag, rng *rand.Rand) {
	for _, post := range posts {
		n := rng.Intn(1)
		db.Model(&post).Association("Tags").Replace(randomSubset(tags, n, rng))
	}
}

func createViews(db *gorm.DB, post database.Post, count int, rng *rand.Rand) {
	for i := 0; i < count; i++ {
		view := database.PostView{
			PostID:    post.ID,
			IPAddress: fmt.Sprintf("192.168.1.%d", rng.Intn(255)),
			UserAgent: "FakeUserAgent/1.0",
			ViewedAt:  time.Now(),
		}
		db.Create(&view)
	}
}

func createComments(db *gorm.DB, post database.Post, count int, authors []database.User, rng *rand.Rand) {
	for i := 0; i < count; i++ {
		author := authors[rng.Intn(len(authors))]
		comment := database.Comment{
			UUID:     uuid.NewString(),
			PostID:   post.ID,
			AuthorID: author.ID,
			Content:  fmt.Sprintf("This is comment #%d", i+1),
		}
		db.Create(&comment)
	}
}

func createLikes(db *gorm.DB, post database.Post, users []database.User, rng *rand.Rand) {
	shuffled := randomSubset(users, 2, rng)
	for _, user := range shuffled {
		like := database.Like{
			PostID: post.ID,
			UserID: user.ID,
		}
		db.Create(&like)
	}
}

func randomSubset[T any](list []T, n int, rng *rand.Rand) []T {
	if n > len(list) {
		n = len(list)
	}

	rng.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})

	return list[:n]
}
