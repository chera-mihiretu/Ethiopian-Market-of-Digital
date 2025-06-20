package entities

import (
	"time"

	"github.com/google/uuid"
)

type SessionEntity struct {
	ID           uuid.UUID `json:"id" bson:"_id"`
	UserID       uuid.UUID `json:"user_id" bson:"user_id"`
	JWTToken     string    `json:"jwt_token" bson:"jwt_token"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at" bson:"expires_at"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}
