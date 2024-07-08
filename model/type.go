package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	Username string
	Password string
}

type Verification struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	UserID           primitive.ObjectID `bson:"user_id"`
	VerifiedAt       time.Time          `bson:"verified_at"`
	IsVerified       bool               `bson:"is_verified"`
	VerificationCode string             `bson:"verification_code"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Salt     string             `bson:"salt,omitempty,omitempty" json:"salt,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
}

type Pengguna struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username     string			    `bson:"username,omitempty" json:"username,omitempty"`
	Akun         User               `bson:"akun,omitempty" json:"akun,omitempty"`
}

type Password struct {
	Password    string `bson:"password,omitempty" json:"password,omitempty"`
	Newpassword string `bson:"newpass,omitempty" json:"newpass,omitempty"`
}

type Payload struct {
	Id   primitive.ObjectID `json:"id"`
	Role string             `json:"role"`
	Exp  time.Time          `json:"exp"`
	Iat  time.Time          `json:"iat"`
	Nbf  time.Time          `json:"nbf"`
}

type Response struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Data    []User `json:"data,omitempty" bson:"data,omitempty"`
	Role    string `json:"role,omitempty" bson:"role,omitempty"`
}

type ChatRequest struct {
    Message string `json:"message"`
}

type ChatResponse struct {
    Response string `json:"response"`
}