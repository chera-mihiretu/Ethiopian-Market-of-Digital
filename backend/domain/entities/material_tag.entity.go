package entities

import (
	"time"

	"github.com/google/uuid"
)

type MaterialTagEntity struct {
	ID         uuid.UUID `json:"id" bson:"_id"`
	MaterialID uuid.UUID `json:"material_id" bson:"material_id"`
	Tag        string    `json:"tag" bson:"tag"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}
