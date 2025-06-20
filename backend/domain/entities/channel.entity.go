package entities

import (
	"time"

	"github.com/google/uuid"
)

type ChannelEntity struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	CreatedBy uuid.UUID `json:"created_by" bson:"created_by"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
