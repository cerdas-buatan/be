package be

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

"github.com/cerdas-buatan/be/model"
"github.com/cerdas-buatan/be/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "cobaiesnn_db")


type Userr struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email string             `bson:"email,omitempty" json:"email,omitempty"`
	Role  string             `bson:"role,omitempty" json:"role,omitempty"`
}

func TestGetAllDoc(t *testing.T) {
	hasil := module.GetAllDocs(db, "user", []Userr{})
	fmt.Println(hasil)
}

func TestInsertUser(t *testing.T) {
	var doc model.User
	doc.Email = "admin@gmail.com"
	password := "admin123"
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		t.Errorf("kesalahan server : salt")
	} else {
		hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
		user := bson.M{
			"email":    doc.Email,
			"password": hex.EncodeToString(hashedPassword),
			"salt":     hex.EncodeToString(salt),
			"role":     "admin",
		}
		_, err = module.InsertOneDoc(db, "user", user)
		if err != nil {
			t.Errorf("gagal insert")
		} else {
			fmt.Println("berhasil insert")
		}
	}
}


func TestSignUpPengguna(t *testing.T) {
	var doc model.Pengguna
	doc.Username = "Sahijatea"
	doc.Akun.Email = "jelemages@gmail.com"
	doc.Akun.Password = "sahijabandung"
	err := module.SignUpPengguna(db, doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Println("Data berhasil disimpan dengan nama :", doc.Username)
	}
}

func TestLogIn(t *testing.T) {
	var doc model.User
	doc.Email = "sahjatsea@gmail.com"
	doc.Password = "sahijabandung"
	user, err := module.LogIn(db, doc)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Selamat datang Driver:", user)
	}
}

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("ini private key :", privateKey)
	fmt.Println("ini public key :", publicKey)
	id := "6569a026a943657839880665"
	objectId, err := primitive.ObjectIDFromHex(id)
	role := "pengguna"
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	hasil, err := module.Encode(objectId, role, privateKey)
	fmt.Println("ini hasil :", hasil, err)
}


func TestWatoken(t *testing.T) {
	body, err := module.Decode("38a1f84dfe0d9bf6678f68646d6bd0fabdd63166f9d173a824d7c7ac9bc0cd4a", "v4.public.eyJleHAiOiIyMDI0LTA3LTA2VDAxOjI4OjAyKzA3OjAwIiwiaWF0IjoiMjAyNC0wNy0wNVQyMzoyODowMiswNzowMCIsImlkIjoiNjU2OWEwMjZhOTQzNjU3ODM5ODgwNjY1IiwibmJmIjoiMjAyNC0wNy0wNVQyMzoyODowMiswNzowMCIsInJvbGUiOiJwZW5nZ3VuYSJ9i8UyAwE4VbAF5GZ4AAyJYZGafEelBeILD2o2Ddcs7Jp0MIY7tZjLjXiwO-GiIGCm72XdEcfASBsuNxDNnqYTAg")
	fmt.Println("isi : ", body, err)
}


func TestReturnStruct(t *testing.T) {
	id := "11b98454e034f3045021a8aa8eb84280"
	objectId, _ := primitive.ObjectIDFromHex(id)
	user, _ := module.GetUserFromID(objectId, db)
	data := model.User{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
	hasil := module.GCFReturnStruct(data)
	fmt.Println(hasil)
}
