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
	Message  string `json:"message"`     
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`  
	Token   string `json:"token,omitempty" bson:"token,omitempty"`  
	Message string `json:"message,omitempty" bson:"message,omitempty"`  
}

type Payload struct {
	ID       string    `json:"id"`
	Role     string    `json:"role"`
	IssuedAt time.Time `json:"issued_at"`
	Expiry   time.Time `json:"expiry"`
}

type Menu struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`  
	Name string             `bson:"name" json:"name"`  
	Link string             `bson:"link" json:"link"`  
}

// delete from here 

// struct Pengguna
type Pengguna struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`  
	Username string             `bson:"username,omitempty" json:"username,omitempty"`  
	Akun     User               `bson:"akun,omitempty" json:"akun,omitempty"`  
}

// struct user
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Salt     string             `bson:"salt,omitempty" json:"salt,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
}