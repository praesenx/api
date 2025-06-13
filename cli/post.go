package main

import (
	"fmt"
	"github.com/oullin/pkg/markdown"
)

func main() {
	file := markdown.Parser{
		Url: "https://raw.githubusercontent.com/oullin/content/refs/heads/main/leadership/2025-04-02-embrace-growth-through-movement.md",
	}

	response, err := file.Fetch()

	if err != nil {
		panic(fmt.Sprintf("Error fetching the markdown content: %v", err))
	}

	post, err := markdown.Parse(response)

	if err != nil {
		panic(fmt.Sprintf("Error parsing markdown: %v", err))
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
}
