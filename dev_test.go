package be

import (
	// "crypto/rand"
	// "encoding/hex"
	"fmt"
	"testing"

	"github.com/serbaevents/backendSE/model"
	"github.com/serbaevents/backendSE/module"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "golang.org/x/crypto/argon2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "serbaevent_db")

type Userr struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email string             `bson:"email,omitempty" json:"email,omitempty"`
	Role  string             `bson:"role,omitempty" json:"role,omitempty"`
}

func TestGetAllDoc(t *testing.T) {
	hasil := module.GetAllDocs(db, "user", []Userr{})
	fmt.Println(hasil)
}
func TestSignUpPengguna(t *testing.T) {
	var doc model.Pengguna
	doc.NamaLengkap = "Sahijatea"
	doc.TanggalLahir = "30/08/2004"
	doc.JenisKelamin = "Perempuan"
	doc.NomorHP = "081234567890"
	doc.Alamat = "Wastukencana Blok No 32"
	doc.Akun.Email = "sahjatea@gmail.com"
	doc.Akun.Password = "sahijabandung"
	err := module.SignUpPengguna(db, doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Println("Data berhasil disimpan dengan nama :", doc.NamaLengkap)
	}
}
