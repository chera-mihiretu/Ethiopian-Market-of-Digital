package entities

import (
	"github.com/google/uuid"
)

type CourseEntity struct {
	ID           uuid.UUID `json:"id" bson:"_id"`
	Name         string    `json:"name" bson:"name"`
	DepartmentID uuid.UUID `json:"department_id" bson:"department_id"`
}
