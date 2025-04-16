package seeds

import (
	"fmt"
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/webkit/gorm"
	"github.com/google/uuid"
)

type CategoriesSeed struct {
	db *database.Connection
}

type CategoriesAttrs struct {
	Slug        string
	Description string
}

func MakeCategoriesSeed(db *database.Connection) *CategoriesSeed {
	return &CategoriesSeed{
		db: db,
	}
}

func (s CategoriesSeed) Create(attrs CategoriesAttrs) ([]database.Category, error) {
	var categories []database.Category

	seeds := []string{
		"Tech", "AI", "Leadership", "Innovation",
		"Cloud", "Data", "DevOps", "ML", "Startups", "Engineering",
	}

	for index, seed := range seeds {
		categories = append(categories, database.Category{
			UUID:        uuid.NewString(),
			Name:        seed,
			Slug:        fmt.Sprintf("[%d]: slug-%s", index+1, attrs.Slug),
			Description: attrs.Description,
		})
	}

	result := s.db.Sql().Create(&categories)

	if gorm.HasDbIssues(result.Error) {
		return nil, fmt.Errorf("error seeding categories: %s", result.Error)
	}

	return categories, nil
}
