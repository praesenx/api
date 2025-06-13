package posts

import (
	"fmt"
	"github.com/oullin/pkg/markdown"
)

func Handle() error {
	uri, err := markdown.ReadURL()

	if err != nil {
		return fmt.Errorf("error reading the URL: %v", err)
	}

	file := markdown.Parser{Url: *uri}

	response, err := file.Fetch()

	if err != nil {
		return fmt.Errorf("error fetching the markdown content: %v", err)
	}

	post, err := markdown.Parse(response)

	if err != nil {
		return fmt.Errorf("error parsing markdown: %v", err)
	}

	// --- All good!
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Excerpt: %s\n", post.Excerpt)
	fmt.Printf("Slug: %s\n", post.Slug)
	fmt.Printf("Author: %s\n", post.Author)
	fmt.Printf("Image URL: %s\n", post.ImageURL)
	fmt.Printf("Image Alt: %s\n", post.ImageAlt)
	fmt.Printf("Category Alt: %s\n", post.Category)
	fmt.Printf("Tags Alt: %s\n", post.Tags)
	fmt.Println("\n--- Content ---")
	fmt.Println(post.Content)

	return nil
}
