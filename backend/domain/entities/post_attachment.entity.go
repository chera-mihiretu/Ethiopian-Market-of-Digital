package entities

import (
	"time"

	"github.com/google/uuid"
)

type PostAttachmentEntity struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	PostID    uuid.UUID `json:"post_id" bson:"post_id"`
	FileURL   string    `json:"file_url" bson:"file_url"`
	FileType  string    `json:"file_type" bson:"file_type"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
