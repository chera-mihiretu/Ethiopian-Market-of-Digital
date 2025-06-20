package entities

import (
	"time"

	"github.com/google/uuid"
)

type ChannelUserEntity struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	ChannelID uuid.UUID `json:"channel_id" bson:"channel_id"`
	UserID    uuid.UUID `json:"user_id" bson:"user_id"`
	Role      string    `json:"role" bson:"role"`
	JoinedAt  time.Time `json:"joined_at" bson:"joined_at"`
}
