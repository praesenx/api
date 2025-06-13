package flow

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io"
    "net/http"
    "regexp"
    "strings"
    "time"
)

// FetchMarkdown downloads the Markdown file from a public URL
func FetchMarkdown(url string) (string, error) {
    // Bust CDN or proxy caches by adding a unique timestamp
    sep := "?"
    if strings.Contains(url, "?") {
        sep = "&"
    }
    timestampedURL := fmt.Sprintf("%s%sts=%d", url, sep, time.Now().UnixNano())

    req, err := http.NewRequest("GET", timestampedURL, nil)
    if err != nil {
        return "", err
    }
    // Instruct intermediate caches to revalidate
    req.Header.Set("Cache-Control", "no-cache")
    req.Header.Set("Pragma", "no-cache")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to fetch markdown: status %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}

// ParseMarkdown splits the document into front-matter and content, then parses YAML
// It also extracts a leading Markdown image (header image) if present
func ParseMarkdown(data string) (Post, error) {
    var post Post
    // Expecting format: ---\n<yaml>---\n<content>
    sections := strings.SplitN(data, "---", 3)
    if len(sections) < 3 {
        return post, fmt.Errorf("invalid front-matter format")
    }

    fm := strings.TrimSpace(sections[1])
    body := strings.TrimSpace(sections[2])

    // Unmarshal YAML into FrontMatter
    err := yaml.Unmarshal([]byte(fm), &post.FrontMatter)
    if err != nil {
        return post, err
    }

    // Look for a header image at the top of the content
    // Markdown image syntax: ![alt](url)
    re := regexp.MustCompile(`^!\[(.*?)\]\((.*?)\)`)

    // Split first line from rest of content
    parts := strings.SplitN(body, "\n", 2)
    first := strings.TrimSpace(parts[0])

    if m := re.FindStringSubmatch(first); len(m) == 3 {
        post.ImageAlt = m[1]
        post.ImageURL = m[2]

        // Remaining content excludes the header image line
        if len(parts) > 1 {
            post.Content = strings.TrimSpace(parts[1])
        } else {
            post.Content = ""
        }
    } else {
        // No header image found; the entire body is content
        post.ImageAlt = ""
        post.ImageURL = ""
        post.Content = body
    }

    return post, nil
}
