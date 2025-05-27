package seeds

import (
	"fmt"
	"github.com/gocanto/blog/database"
	"github.com/google/uuid"
	"time"
)

type CommentsSeed struct {
	db *database.Connection
}

type CommentsAttrs struct {
	UUID       string
	PostID     uint64
	AuthorID   uint64
	ParentID   *uint64
	Content    string
	ApprovedAt *time.Time
}

func MakeCommentsSeed(db *database.Connection) *CommentsSeed {
	return &CommentsSeed{
		db: db,
	}
}

func (s CommentsSeed) Create(attrs ...CommentsAttrs) ([]database.Comment, error) {
	var comments []database.Comment

	for _, attr := range attrs {
		comments = append(comments, database.Comment{
			UUID:       uuid.NewString(),
			PostID:     attr.PostID,
			AuthorID:   attr.AuthorID,
			ParentID:   attr.ParentID,
			Content:    attr.Content,
			ApprovedAt: attr.ApprovedAt,
		})
	}

	result := s.db.Sql().Create(&comments)

	if result.Error != nil {
		return nil, fmt.Errorf("error creating comments: %w", result.Error)
	}

	// ---- Parent Updates
	firstComment := comments[0]
	lastComment := comments[len(comments)-1]

	result = s.db.Sql().Model(database.Comment{}).
		Where("id = ?", lastComment.ID).
		Update("parent_id", firstComment.ID)

	if result.Error != nil {
		return nil, fmt.Errorf("error updating for parent comment: %w", result.Error)
	}

	return comments, nil
}
