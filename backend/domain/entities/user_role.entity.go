package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserRoleEntity struct {
	ID         uuid.UUID `json:"id" bson:"_id"`
	UserID     uuid.UUID `json:"user_id" bson:"user_id"`
	RoleID     uuid.UUID `json:"role_id" bson:"role_id"`
	AssignedAt time.Time `json:"assigned_at" bson:"assigned_at"`
}
