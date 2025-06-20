package entities

import (
	"time"

	"github.com/google/uuid"
)

type WhatsNewEntity struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
