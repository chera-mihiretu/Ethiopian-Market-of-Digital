package entities

import (
	"time"

	"github.com/google/uuid"
)

type OpportunityApplicationEntity struct {
	ID            uuid.UUID `json:"id" bson:"_id"`
	OpportunityID uuid.UUID `json:"opportunity_id" bson:"opportunity_id"`
	ApplicantID   uuid.UUID `json:"applicant_id" bson:"applicant_id"`
	ResumeLink    string    `json:"resume_link" bson:"resume_link"`
	AppliedAt     time.Time `json:"applied_at" bson:"applied_at"`
}
