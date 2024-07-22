package module

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/argon2"

//	"github.com/badoux/checkmail"
//	"golang.org/x/crypto/argon2"

	model "github.com/cerdas-buatan/be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


// signup
func SignUpPengguna(db *mongo.Database, insertedDoc model.Pengguna) error {
	objectId := primitive.NewObjectID()

		// Validate mandatory fields
	if insertedDoc.Username == "" || insertedDoc.Akun.Password == "" {
		return fmt.Errorf("dimohon untuk melengkapi data")
	}

	// Validate email format
	if err := checkmail.ValidateFormat(insertedDoc.Akun.Email); err != nil {
		return fmt.Errorf("email tidak valid")
	}

	// Check if the email is already registered
	userExists, _ := GetUserFromEmail(insertedDoc.Akun.Email, db)
	if userExists.Email != "" {
		return fmt.Errorf("email sudah terdaftar")
	}

	// Validate password constraints
	if strings.Contains(insertedDoc.Akun.Password, " ") {
		return fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Akun.Password) < 8 {
		return fmt.Errorf("password terlalu pendek")
	}


	// Generate salt and hash the password
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("kesalahan server: gagal membuat salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Akun.Password), salt, 1, 64*1024, 4, 32)


	// Create user and pengguna documents
	user := bson.M{
		"_id":      objectId,
		"email":    insertedDoc.Akun.Email,
		"password": hex.EncodeToString(hashedPassword),
		"salt":     hex.EncodeToString(salt),
	}
	pengguna := bson.M{
		"username": insertedDoc.Username,
		"akun": model.User{
			ID: objectId,
		},
	}

	
	// Insert the documents into the database
	if _, err := InsertOneDoc(db, "user", user); err != nil {
		return fmt.Errorf("kesalahan server: gagal menyimpan user")
	}
	if _, err := InsertOneDoc(db, "pengguna", pengguna); err != nil {
		return fmt.Errorf("kesalahan server: gagal menyimpan pengguna")
	}

	return nil
}

// login
func LogIn(db *mongo.Database, insertedDoc model.User) (user model.User, err error) {
    // Validate mandatory fields
    if insertedDoc.Email == "" || insertedDoc.Password == "" {
        return user, fmt.Errorf("mohon untuk melengkapi data")
    }

    // Validate email format
    if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
        return user, fmt.Errorf("email tidak valid")
    }

    // Retrieve the user from the database
    existsDoc, err := GetUserFromEmail(insertedDoc.Email, db)
    if err != nil {
        return user, fmt.Errorf("kesalahan server : gagal mengambil data pengguna")
    }

    // Decode the stored salt
    salt, err := hex.DecodeString(existsDoc.Salt)
    if err != nil {
        return user, fmt.Errorf("kesalahan server : salt tidak valid")
    }

    // Hash the provided password with the stored salt
    hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)

    // Compare the hashed password with the stored password
    if hex.EncodeToString(hash) != existsDoc.Password {
        return user, fmt.Errorf("password salah")
    }

    // Successful login, return the existing user document
    return existsDoc, nil
}

func GetUserFromEmail(email string, db *mongo.Database) (doc model.User, err error) {
	// Mendapatkan koleksi "user" dari database
	collection := db.Collection("user")
	
	// Membuat filter untuk pencarian berdasarkan email
	filter := bson.M{"email": email}
	
	// Mencari satu dokumen yang sesuai dengan filter
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	
	// Memeriksa apakah terjadi error
	if err != nil {
		// Jika tidak ada dokumen yang ditemukan, mengembalikan error "email tidak ditemukan"
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("email tidak ditemukan")
		}
		// Jika terjadi kesalahan lain, mengembalikan error "kesalahan server"
		return doc, fmt.Errorf("kesalahan server")
	}
	
	// Mengembalikan dokumen pengguna jika ditemukan
	return doc, nil
}

func GetAllUser(db *mongo.Database) (user []model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return user, fmt.Errorf("error GetAllUser mongo: %s", err)
	}
	err = cursor.All(context.Background(), &user)
	if err != nil {
		return user, fmt.Errorf("error GetAllUser context: %s", err)
	}
	return user, nil
}

func GetPenggunaFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.Pengguna, err error) {
	// Mendapatkan koleksi "pengguna" dari database
	collection := db.Collection("pengguna")
	
	// Membuat filter untuk pencarian berdasarkan ID
	filter := bson.M{"_id": _id}
	
	// Mencari satu dokumen yang sesuai dengan filter
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	
	// Memeriksa apakah terjadi error
	if err != nil {
		// Jika tidak ada dokumen yang ditemukan, mengembalikan error "tidak ada data untuk ID <id>"
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("tidak ada data untuk ID %s", _id.Hex())
		}
		// Jika terjadi kesalahan lain, mengembalikan error "kesalahan saat mengambil data untuk ID <id>: <pesan error>"
		return doc, fmt.Errorf("kesalahan saat mengambil data untuk ID %s: %s", _id.Hex(), err.Error())
	}
	
	// Mengembalikan dokumen pengguna jika ditemukan
	return doc, nil
}

// GetResponse generates a response based on the input message
func GCFGetResponse(message string, db *mongo.Database) (string, error) {
	collection := db.Collection("chat_responses")

	var Responsechat model.ChatResponse
	filter := bson.M{"message": message}

	err := collection.FindOne(context.TODO(), filter).Decode(&Responsechat)
	if err != nil {
		return "", err
	}

	return Responsechat.Response, nil
}


// insert user
func InsertUser(db *mongo.Database, collection string, userdata User) string {
	hash, _ := HashPassword(userdata.Password)
	userdata.Password = hash
	atdb.InsertOneDoc(db, collection, userdata)
	return "username : " + userdata.Username + "password : " + userdata.Password
}

func GCFPredict(w http.ResponseWriter, r *http.Request) {
	predictHandler(w, r)
}

//ChatHandler
func ChatHandler(w http.ResponseWriter, r *http.Request) {
    var chatReq model.ChatRequest
    err := json.NewDecoder(r.Body).Decode(&chatReq)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response, err := module.GetResponse(chatReq.Message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    chatRes := model.ChatResponse{Response: response}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(chatRes)
}
