package entities

import (
	"time"

	"github.com/google/uuid"
)

type MaterialEntity struct {
	ID               uuid.UUID  `json:"id" bson:"_id"`
	Title            string     `json:"title" bson:"title"`
	DepartmentID     uuid.UUID  `json:"department_id" bson:"department_id"`
	CourseID         uuid.UUID  `json:"course_id" bson:"course_id"`
	Year             int        `json:"year" bson:"year"`
	FileURL          string     `json:"file_url" bson:"file_url"`
	UploadedBy       uuid.UUID  `json:"uploaded_by" bson:"uploaded_by"`
	ParentMaterialID *uuid.UUID `json:"parent_material_id,omitempty" bson:"parent_material_id,omitempty"`
	IsVerified       bool       `json:"is_verified" bson:"is_verified"`
	CreatedAt        time.Time  `json:"created_at" bson:"created_at"`
}
