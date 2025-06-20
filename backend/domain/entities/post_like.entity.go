package entities

import (
	"time"

	"github.com/google/uuid"
)

type PostLikeEntity struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	PostID    uuid.UUID `json:"post_id" bson:"post_id"`
	UserID    uuid.UUID `json:"user_id" bson:"user_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
