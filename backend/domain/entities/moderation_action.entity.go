package entities

import (
	"time"

	"github.com/google/uuid"
)

type ModerationActionEntity struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	ReportID  uuid.UUID `json:"report_id" bson:"report_id"`
	AdminID   uuid.UUID `json:"admin_id" bson:"admin_id"`
	Action    string    `json:"action" bson:"action"`
	Notes     string    `json:"notes" bson:"notes"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
