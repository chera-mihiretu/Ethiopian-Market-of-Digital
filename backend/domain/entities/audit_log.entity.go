package entities

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditLogEntity struct {
	ID         uuid.UUID   `json:"id" bson:"_id"`
	UserID     uuid.UUID   `json:"user_id" bson:"user_id"`
	Action     string      `json:"action" bson:"action"`
	TargetType string      `json:"target_type" bson:"target_type"`
	TargetID   uuid.UUID   `json:"target_id" bson:"target_id"`
	Details    primitive.M `json:"details" bson:"details"`
	CreatedAt  time.Time   `json:"created_at" bson:"created_at"`
}
