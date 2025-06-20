package entities

import (
	"time"

	"github.com/google/uuid"
)

type ReportEntity struct {
	ID             uuid.UUID `json:"id" bson:"_id"`
	ReportedBy     uuid.UUID `json:"reported_by" bson:"reported_by"`
	ReportedPostID uuid.UUID `json:"reported_post_id" bson:"reported_post_id"`
	Reason         string    `json:"reason" bson:"reason"`
	Reviewed       bool      `json:"reviewed" bson:"reviewed"`
	ReviewedBy     uuid.UUID `json:"reviewed_by" bson:"reviewed_by"`
	ReviewedAt     time.Time `json:"reviewed_at" bson:"reviewed_at"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
}
