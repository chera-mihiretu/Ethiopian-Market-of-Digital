package entities

import (
	"time"

	"github.com/google/uuid"
)

type CommentEntity struct {
	ID              uuid.UUID  `json:"id" bson:"_id"`
	PostID          uuid.UUID  `json:"post_id" bson:"post_id"`
	UserID          uuid.UUID  `json:"user_id" bson:"user_id"`
	ParentCommentID *uuid.UUID `json:"parent_comment_id,omitempty" bson:"parent_comment_id,omitempty"`
	Content         string     `json:"content" bson:"content"`
	CreatedAt       time.Time  `json:"created_at" bson:"created_at"`
}
