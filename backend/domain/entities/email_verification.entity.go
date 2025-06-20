package entities

import (
	"time"

	"github.com/google/uuid"
)

type EmailVerificationEntity struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	UserID    uuid.UUID `json:"user_id" bson:"user_id"`
	Token     string    `json:"token" bson:"token"`
	SentAt    time.Time `json:"sent_at" bson:"sent_at"`
	ExpiresAt time.Time `json:"expires_at" bson:"expires_at"`
	Consumed  bool      `json:"consumed" bson:"consumed"`
}
