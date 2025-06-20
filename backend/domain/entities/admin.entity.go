package entities

import (
	"time"

	"github.com/google/uuid"
)

type AdminEntity struct {
	ID           uuid.UUID `json:"id" bson:"_id"`
	UserID       uuid.UUID `json:"user_id" bson:"user_id"`
	IsSuperAdmin bool      `json:"is_super_admin" bson:"is_super_admin"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}
