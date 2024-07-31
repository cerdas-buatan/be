package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// struct user
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Salt     string             `bson:"salt,omitempty" json:"salt,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
}

// struct Pengguna
type Pengguna struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Akun     User               `bson:"akun,omitempty" json:"akun,omitempty"`
}

// struct response
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// struct ChatRequest
type ChatRequest struct {
	Message string `json:"message"`
}

// ChatResponse represents a chat response
type ChatResponse struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Message  string             `bson:"message" json:"message"`
	Response string             `bson:"response" json:"response"`
}

type Credential struct {
//	ID       string    `json:"id"`
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Payload struct {
	ID       string    `json:"id"`
	// Role     string    `json:"role"`
	IssuedAt time.Time `json:"issued_at"`
	Expiry   time.Time `json:"expiry"`
}

type Menu struct {
    ID   primitive.ObjectID `bson:"_id,omitempty"`
    Name string             `bson:"name"`

}

type MenuService struct {
    collection *mongo.Collection
}

type ChatHistory struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    UserID    string             `bson:"user_id"`
    Timestamp string             `bson:"timestamp"`
    Message   string             `bson:"message"`
}

type Chats struct {
	IDChats   string `json:"idchats"`
	Message   string `json:"message"`
	Responses string `json:"responses"`
	Score     float64
}

type Dataset struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Question string             `bson:"question" json:"question"`
	Answer   string             `bson:"answer" json:"answer"`
}

// struct ForgotPasswordRequest
type ForgotPasswordRequest struct {
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type VerificationCode struct {
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Code        string `json:"code" bson:"code"`
	ExpiresAt   int64  `json:"expires_at" bson:"expires_at"`
}
