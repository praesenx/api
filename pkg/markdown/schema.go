package markdown

type FrontMatter struct {
	Title    string   `yaml:"title"`
	Excerpt  string   `yaml:"excerpt"`
	Slug     string   `yaml:"slug"`
	Author   string   `yaml:"author"`
	Category string   `yaml:"category"`
	Tags     []string `yaml:"tags"`
}

type Post struct {
	FrontMatter
	ImageURL string
	ImageAlt string
	Content  string
}

type File struct {
	Url string
}
