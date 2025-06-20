package entities

import (
	"time"

	"github.com/google/uuid"
)

type FollowEntity struct {
	ID         uuid.UUID `json:"id" bson:"_id"`
	FollowerID uuid.UUID `json:"follower_id" bson:"follower_id"`
	FollowedID uuid.UUID `json:"followed_id" bson:"followed_id"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}
