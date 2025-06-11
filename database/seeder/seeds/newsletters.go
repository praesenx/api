package seeds

import (
	"fmt"
	"github.com/oullin/database"
	"github.com/oullin/pkg/gorm"
	"time"
)

type NewslettersSeed struct {
	db *database.Connection
}

type NewsletterAttrs struct {
	FirstName      string
	LastName       string
	Email          string
	SubscribedAt   *time.Time
	UnsubscribedAt *time.Time
}

func MakeNewslettersSeed(db *database.Connection) *NewslettersSeed {
	return &NewslettersSeed{
		db: db,
	}
}

func (s NewslettersSeed) Create(attrs []NewsletterAttrs) error {
	var newsletters []database.Newsletter

	for _, attr := range attrs {
		letter := database.Newsletter{
			FirstName:      attr.FirstName,
			LastName:       attr.LastName,
			Email:          attr.Email,
			SubscribedAt:   attr.SubscribedAt,
			UnsubscribedAt: attr.UnsubscribedAt,
		}

		newsletters = append(newsletters, letter)
	}

	result := s.db.Sql().Create(&newsletters)

	if gorm.HasDbIssues(result.Error) {
		return fmt.Errorf("error seeding newsletters: %s", result.Error)
	}

	return nil
}
