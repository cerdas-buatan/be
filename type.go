package type

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct user
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Salt     string             `bson:"salt,omitempty" json:"salt,omitempty"`
	CreatedAt time.Time         `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time         `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}


// struct Pengguna
type Pengguna struct {
	Username string             `bson:"username,omitempty" json:"username,omitempty"`  
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
	Message  string `json:"message"`      
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`  
	Token   string `json:"token,omitempty" bson:"token,omitempty"`    
	Message string `json:"message,omitempty" bson:"message,omitempty"`    
}  

type ChatResponse struct {
	ResponseID primitive.ObjectID `bson:"response_id,omitempty" json:"response_id,omitempty"`
	Message    string             `json:"message"`   
	Timestamp  int64              `json:"timestamp"` 
	UserID     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
}
  
// struct LoginRequest
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// struct LoginResponse
type LoginResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	UserID  primitive.ObjectID `json:"user_id,omitempty"`
}
  
// struct RegisterRequest
type RegisterRequest struct {
	Username string `json:"username"` 
	Email    string `json:"email"` 
	Password string `json:"password"` 
} 

// struct RegisterResponse
type RegisterResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	UserID  primitive.ObjectID `json:"user_id,omitempty"`
}