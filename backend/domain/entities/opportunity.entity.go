package entities

import (
	"time"

	"github.com/google/uuid"
)

type OpportunityEntity struct {
	ID           uuid.UUID `json:"id" bson:"_id"`
	DepartmentID uuid.UUID `json:"department_id" bson:"department_id"`
	Title        string    `json:"title" bson:"title"`
	Description  string    `json:"description" bson:"description"`
	Link         string    `json:"link" bson:"link"`
	Type         string    `json:"type" bson:"type"`
	PostedBy     uuid.UUID `json:"posted_by" bson:"posted_by"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}
