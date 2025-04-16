package seeds

import (
	"fmt"
	"github.com/gocanto/blog/database"
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

func (s TagsSeed) Create() []database.Tag {
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

	s.db.Sql().Create(&tags)

	return tags
}
