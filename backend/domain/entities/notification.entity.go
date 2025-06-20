package entities

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationEntity struct {
	ID        uuid.UUID   `json:"id" bson:"_id"`
	UserID    uuid.UUID   `json:"user_id" bson:"user_id"`
	Type      string      `json:"type" bson:"type"`
	Payload   primitive.M `json:"payload" bson:"payload"`
	IsRead    bool        `json:"is_read" bson:"is_read"`
	CreatedAt time.Time   `json:"created_at" bson:"created_at"`
}
