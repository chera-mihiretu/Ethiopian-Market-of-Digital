package entities

import (
	"time"

	"github.com/google/uuid"
)

type PostEntity struct {
	ID             uuid.UUID `json:"id" bson:"_id"`
	UserID         uuid.UUID `json:"user_id" bson:"user_id"`
	Content        string    `json:"content" bson:"content"`
	IsAnnouncement bool      `json:"is_announcement" bson:"is_announcement"`
	IsValidated    bool      `json:"is_validated" bson:"is_validated"`
	IsFlagged      bool      `json:"is_flagged" bson:"is_flagged"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}
