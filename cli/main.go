package main

import (
    "fmt"
    "github.com/oullin/pkg/flow"
    "log"
)

func main() {
    url := "https://raw.githubusercontent.com/oullin/content/refs/heads/main/leadership/2025-04-02-embrace-growth-through-movement.md"

    data, err := flow.FetchMarkdown(url)
    if err != nil {
        log.Fatalf("Error fetching markdown: %v", err)
    }

    post, err := flow.ParseMarkdown(data)
    if err != nil {
        log.Fatalf("Error parsing markdown: %v", err)
    }

    // Output parsed fields
    fmt.Printf("Title: %s\n", post.Title)
    fmt.Printf("Excerpt: %s\n", post.Excerpt)
    fmt.Printf("Slug: %s\n", post.Slug)
    fmt.Printf("Author: %s\n", post.Author)
    fmt.Printf("Image URL: %s\n", post.ImageURL)
    fmt.Printf("Image Alt: %s\n", post.ImageAlt)
    fmt.Printf("Category Alt: %s\n", post.Category)
    fmt.Printf("Tags Alt: %s\n", post.Tags)
    fmt.Println("--- Content ---")
    fmt.Println(post.Content)
}
