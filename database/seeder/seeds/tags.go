package seeds

import (
	"fmt"
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/webkit/gorm"
	"github.com/google/uuid"
)

type TagsSeed struct {
	db *database.Connection
}

func MakeTagsSeed(db *database.Connection) *TagsSeed {
	return &TagsSeed{
		db: db,
	}
}

func (s TagsSeed) Create() ([]database.Tag, error) {
	var tags []database.Tag
	allowed := []string{
		"Tech", "AI", "Leadership", "Ethics",
		"Automation", "Teamwork", "Agile", "OpenAI", "Scaling", "Future",
	}

	for index, name := range allowed {
		tag := database.Tag{
			UUID: uuid.NewString(),
			Name: name,
			Slug: fmt.Sprintf("tag[%d]-slug-%s", index, name),
		}

		tags = append(tags, tag)
	}

	result := s.db.Sql().Create(&tags)

	if gorm.HasDbIssues(result.Error) {
		return nil, fmt.Errorf("issues creating tags: %s", result.Error)
	}

	return tags, nil
}
