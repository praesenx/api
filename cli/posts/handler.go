package posts

import (
	"fmt"
	"github.com/oullin/pkg/cli"
	"github.com/oullin/pkg/markdown"
)

func (i *Input) Parse() error {
	file := markdown.Parser{
		Url: i.Url,
	}

	response, err := file.Fetch()

	if err != nil {
		return fmt.Errorf("%sError fetching the markdown content: %v %s", cli.Red, err, cli.Reset)
	}

	post, err := markdown.Parse(response)

	if err != nil {
		return fmt.Errorf("%sEerror parsing markdown: %v %s", cli.Red, err, cli.Reset)
	}

	// --- All good!
	// Todo: Save post in the DB.
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
