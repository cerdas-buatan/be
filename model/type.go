package model

import (
    "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Register struct {
    Email           string
    Username        string
    Password        string
    ConfirmPassword string
}

type Login struct {
    Username    string
    Password    string
}

type Verification struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	VerifiedAt  time.Time          `bson:"verified_at"`
	IsVerified  bool               `bson:"is_verified"`
	VerificationCode string       `bson:"verification_code"`
}

