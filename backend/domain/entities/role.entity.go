package entities

import (
	"github.com/google/uuid"
)

type RoleEntity struct {
	ID   uuid.UUID `json:"id" bson:"_id"`
	Name string    `json:"name" bson:"name"`
}
