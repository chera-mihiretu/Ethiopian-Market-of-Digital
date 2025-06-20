package entities

import (
	"time"

	"github.com/google/uuid"
)

type CommentLikeEntity struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	CommentID uuid.UUID `json:"comment_id" bson:"comment_id"`
	UserID    uuid.UUID `json:"user_id" bson:"user_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
