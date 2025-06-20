package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID              uuid.UUID `json:"id" bson:"_id"`
	Name            string    `json:"name" bson:"name"`
	Email           string    `json:"email" bson:"email"`
	PasswordHash    string    `json:"password_hash" bson:"password_hash"`
	ProfileImageURL string    `json:"profile_image_url" bson:"profile_image_url"`
	IsVerified      bool      `json:"is_verified" bson:"is_verified"`
	IsTeacher       bool      `json:"is_teacher" bson:"is_teacher"`
	BlueBadge       bool      `json:"blue_badge" bson:"blue_badge"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`
}
