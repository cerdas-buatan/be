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

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Salt     string             `bson:"salt,omitempty,omitempty" json:"salt,omitempty"`
	Role     string             `bson:"role,omitempty" json:"role,omitempty"`
}

type Pengguna struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap  string             `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	TanggalLahir string             `bson:"tanggallahir,omitempty" json:"tanggallahir,omitempty"`
	JenisKelamin string             `bson:"jeniskelamin,omitempty" json:"jeniskelamin,omitempty"`
	NomorHP      string             `bson:"nomorhp,omitempty" json:"nomorhp,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Akun         User               `bson:"akun,omitempty" json:"akun,omitempty"`
}