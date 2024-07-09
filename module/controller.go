// package module

// import (
// 	"context"
// 	"crypto/rand"
// 	"encoding/hex"

// 	// "encoding/json"
// 	"errors"
// 	"fmt"

// 	// "net/http"
// 	"os"
// 	"os/exec"
// 	"strings"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"golang.org/x/crypto/argon2"

// 	model "github.com/cerdas-buatan/be/model"
// )

// //register

// func RegisterUser(c *fiber.Ctx) error {
// 	db := c.Locals("db").(*mongo.Database)

// 	username := c.FormValue("username")
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")

// 	passwordHash, err := helper.HashPassword(password)
// 	if err != nil {
// 		return helper.SendResponse(c, fiber.StatusInternalServerError, "Failed to hash password", nil)
// 	}

// 	user := model.User{
// 		Username:     username,
// 		Email:        email,
// 		Password:     password,
// 		PasswordHash: passwordHash,
// 	}

// 	_, err = helper.InsertOneDoc(db, "users", user)
// 	if err != nil {
// 		return helper.SendResponse(c, fiber.StatusInternalServerError, "Failed to insert user", nil)
// 	}

// 	return helper.SendResponse(c, fiber.StatusCreated, "User registered successfully", user)
// }

// // signup
// func SignUpPengguna(db *mongo.Database, insertedDoc model.Pengguna) error {
// 	objectId := primitive.NewObjectID()
// 	if insertedDoc.Username == "" ||
// 		insertedDoc.Akun.Password == "" {
// 		return fmt.Errorf("dimohon untuk melengkapi data")
// 	}
// 	if err := checkmail.ValidateFormat(insertedDoc.Akun.Email); err != nil {
// 		return fmt.Errorf("email tidak valid")
// 	}
// 	userExists, _ := GetUserFromEmail(insertedDoc.Akun.Email, db)
// 	if insertedDoc.Akun.Email == userExists.Email {
// 		return fmt.Errorf("email sudah terdaftar")
// 	}
// 	if strings.Contains(insertedDoc.Akun.Password, " ") {
// 		return fmt.Errorf("password tidak boleh mengandung spasi")
// 	}
// 	if len(insertedDoc.Akun.Password) < 8 {
// 		return fmt.Errorf("password terlalu pendek")
// 	}
// 	salt := make([]byte, 16)
// 	_, err := rand.Read(salt)
// 	if err != nil {
// 		return fmt.Errorf("kesalahan server : salt")
// 	}
// 	hashedPassword := argon2.IDKey([]byte(insertedDoc.Akun.Password), salt, 1, 64*1024, 4, 32)
// 	user := bson.M{
// 		"_id":      objectId,
// 		"email":    insertedDoc.Akun.Email,
// 		"password": hex.EncodeToString(hashedPassword),
// 		"salt":     hex.EncodeToString(salt),
// 		"role":     "pengguna",
// 	}
// 	pengguna := bson.M{
// 		"username": insertedDoc.Username,
// 		"akun": model.User{
// 			ID: objectId,
// 		},
// 	}
// 	_, err = InsertOneDoc(db, "user", user)
// 	if err != nil {
// 		return fmt.Errorf("kesalahan server")
// 	}
// 	_, err = InsertOneDoc(db, "pengguna", pengguna)
// 	if err != nil {
// 		return fmt.Errorf("kesalahan server")
// 	}
// 	return nil
// }

// // login
// func LogIn(db *mongo.Database, insertedDoc model.User) (user model.User, err error) {
// 	if insertedDoc.Email == "" || insertedDoc.Password == "" {
// 		return user, fmt.Errorf("mohon untuk melengkapi data")
// 	}
// 	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
// 		return user, fmt.Errorf("email tidak valid")
// 	}
// 	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db)
// 	if err != nil {
// 		return
// 	}
// 	salt, err := hex.DecodeString(existsDoc.Salt)
// 	if err != nil {
// 		return user, fmt.Errorf("kesalahan server : salt")
// 	}
// 	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
// 	if hex.EncodeToString(hash) != existsDoc.Password {
// 		return user, fmt.Errorf("password salah")
// 	}
// 	return existsDoc, nil
// }

// // user
// func UpdateEmailUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.User) error {
// 	dataUser, err := GetUserFromID(iduser, db)
// 	if err != nil {
// 		return err
// 	}
// 	if insertedDoc.Email == "" {
// 		return fmt.Errorf("dimohon untuk melengkapi data")
// 	}
// 	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
// 		return fmt.Errorf("email tidak valid")
// 	}
// 	existsDoc, _ := GetUserFromEmail(insertedDoc.Email, db)
// 	if existsDoc.Email == insertedDoc.Email {
// 		return fmt.Errorf("email sudah terdaftar")
// 	}
// 	user := bson.M{
// 		"email":    insertedDoc.Email,
// 		"password": dataUser.Password,
// 		"salt":     dataUser.Salt,
// 		"role":     dataUser.Role,
// 	}
// 	err = UpdateOneDoc(iduser, db, "user", user)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //responsechat
// func GetResponse(input string) (string, error) {
//     cmd := exec.Command("python3", "model/model.py", input)
//     out, err := cmd.Output()
//     if err != nil {
//         return "", err
//     }
//     return string(out), nil
// }

// func UpdatePasswordUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.Password) error {
// 	dataUser, err := GetUserFromID(iduser, db)
// 	if err != nil {
// 		return err
// 	}
// 	salt, err := hex.DecodeString(dataUser.Salt)
// 	if err != nil {
// 		return fmt.Errorf("kesalahan server : salt")
// 	}
// 	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
// 	if hex.EncodeToString(hash) != dataUser.Password {
// 		return fmt.Errorf("password lama salah")
// 	}
// 	if insertedDoc.Newpassword == "" {
// 		return fmt.Errorf("mohon untuk melengkapi data")
// 	}
// 	if strings.Contains(insertedDoc.Newpassword, " ") {
// 		return fmt.Errorf("password tidak boleh mengandung spasi")
// 	}
// 	if len(insertedDoc.Newpassword) < 8 {
// 		return fmt.Errorf("password terlalu pendek")
// 	}
// 	salt = make([]byte, 16)
// 	_, err = rand.Read(salt)
// 	if err != nil {
// 		return fmt.Errorf("kesalahan server : salt")
// 	}
// 	hashedPassword := argon2.IDKey([]byte(insertedDoc.Newpassword), salt, 1, 64*1024, 4, 32)
// 	user := bson.M{
// 		"email":    dataUser.Email,
// 		"password": hex.EncodeToString(hashedPassword),
// 		"salt":     hex.EncodeToString(salt),
// 		"role":     dataUser.Role,
// 	}
// 	err = UpdateOneDoc(iduser, db, "user", user)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UpdateUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.User) error {
// 	dataUser, err := GetUserFromID(iduser, db)
// 	if err != nil {
// 		return err
// 	}
// 	if insertedDoc.Email == "" || insertedDoc.Password == "" {
// 		return fmt.Errorf("mohon untuk melengkapi data")
// 	}
// 	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
// 		return fmt.Errorf("email tidak valid")
// 	}
// 	existsDoc, _ := GetUserFromEmail(insertedDoc.Email, db)
// 	if existsDoc.Email == insertedDoc.Email {
// 		return fmt.Errorf("email sudah terdaftar")
// 	}
// 	if strings.Contains(insertedDoc.Password, " ") {
// 		return fmt.Errorf("password tidak boleh mengandung spasi")
// 	}
// 	if len(insertedDoc.Password) < 8 {
// 		return fmt.Errorf("password terlalu pendek")
// 	}
// 	salt := make([]byte, 16)
// 	_, err = rand.Read(salt)
// 	if err != nil {
// 		return fmt.Errorf("kesalahan server : salt")
// 	}
// 	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
// 	user := bson.M{
// 		"email":    insertedDoc.Email,
// 		"password": hex.EncodeToString(hashedPassword),
// 		"salt":     hex.EncodeToString(salt),
// 		"role":     dataUser.Role,
// 	}
// 	err = UpdateOneDoc(iduser, db, "user", user)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetAllUser(db *mongo.Database) (user []model.User, err error) {
// 	collection := db.Collection("user")
// 	filter := bson.M{}
// 	cursor, err := collection.Find(context.Background(), filter)
// 	if err != nil {
// 		return user, fmt.Errorf("error GetAllUser mongo: %s", err)
// 	}
// 	err = cursor.All(context.Background(), &user)
// 	if err != nil {
// 		return user, fmt.Errorf("error GetAllUser context: %s", err)
// 	}
// 	return user, nil
// }

// func GetUserFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.User, err error) {
// 	collection := db.Collection("user")
// 	filter := bson.M{"_id": _id}
// 	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return doc, fmt.Errorf("no data found for ID %s", _id)
// 		}
// 		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
// 	}
// 	return doc, nil
// }

// func GetUserFromEmail(email string, db *mongo.Database) (doc model.User, err error) {
// 	collection := db.Collection("user")
// 	filter := bson.M{"email": email}
// 	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return doc, fmt.Errorf("email tidak ditemukan")
// 		}
// 		return doc, fmt.Errorf("kesalahan server")
// 	}
// 	return doc, nil
// }

// // pengguna
// func UpdatePengguna(idparam, iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.Pengguna) error {
// 	pengguna, err := GetPenggunaFromAkun(iduser, db)
// 	if err != nil {
// 		return err
// 	}
// 	if pengguna.ID != idparam {
// 		return fmt.Errorf("anda bukan pemilik data ini")
// 	}
// 	if insertedDoc.Username == "" {
// 		return fmt.Errorf("dimohon untuk melengkapi data")
// 	}
// 	pgn := bson.M{
// 		"username": insertedDoc.Username,
// 		"akun": model.User{
// 			ID: pengguna.Akun.ID,
// 		},
// 	}
// 	err = UpdateOneDoc(idparam, db, "pengguna", pgn)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetAllPengguna(db *mongo.Database) (pengguna []model.Pengguna, err error) {
// 	collection := db.Collection("pengguna")
// 	filter := bson.M{}
// 	cursor, err := collection.Find(context.Background(), filter)
// 	if err != nil {
// 		return pengguna, fmt.Errorf("error GetAllPengguna mongo: %s", err)
// 	}
// 	err = cursor.All(context.Background(), &pengguna)
// 	if err != nil {
// 		return pengguna, fmt.Errorf("error GetAllPengguna context: %s", err)
// 	}
// 	return pengguna, nil
// }

// func GetPenggunaFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.Pengguna, err error) {
// 	collection := db.Collection("pengguna")
// 	filter := bson.M{"_id": _id}
// 	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return doc, fmt.Errorf("no data found for ID %s", _id)
// 		}
// 		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
// 	}
// 	return doc, nil
// }

// func GetPenggunaFromAkun(akun primitive.ObjectID, db *mongo.Database) (doc model.Pengguna, err error) {
// 	collection := db.Collection("pengguna")
// 	filter := bson.M{"akun._id": akun}
// 	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return doc, fmt.Errorf("pengguna tidak ditemukan")
// 		}
// 		return doc, fmt.Errorf("kesalahan server")
// 	}
// 	return doc, nil
// }


package module

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/cerdas-buatan/be/helper"
	"github.com/cerdas-buatan/be/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// RegisterUser handles user registration
func RegisterUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*mongo.Database)

	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return helper.SendResponse(c, fiber.StatusBadRequest, "Failed to parse request body", nil)
	}

	passwordHash, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.SendResponse(c, fiber.StatusInternalServerError, "Failed to hash password", nil)
	}

	user.Password = passwordHash

	_, err = helper.InsertOneDoc(db, "users", user)
	if err != nil {
		return helper.SendResponse(c, fiber.StatusInternalServerError, "Failed to insert user", nil)
	}

	return helper.SendResponse(c, fiber.StatusCreated, "User registered successfully", user)
}

// Login handles user login
func Login(c *fiber.Ctx) error {
	db := c.Locals("db").(*mongo.Database)

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&credentials); err != nil {
		return helper.SendResponse(c, fiber.StatusBadRequest, "Failed to parse request body", nil)
	}

	user := model.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	authenticatedUser, err := helper.LogIn(db, user)
	if err != nil {
		return helper.SendResponse(c, fiber.StatusUnauthorized, "Authentication failed", nil)
	}

	// Here you can optionally generate a token or session for the user if needed
	return helper.SendResponse(c, fiber.StatusOK, "Login successful", authenticatedUser)
}

// Logout handles user logout (example implementation)
func Logout(c *fiber.Ctx) error {
	// Perform logout logic here, such as clearing session or token
	return helper.SendResponse(c, fiber.StatusOK, "Logout successful", nil)
}
