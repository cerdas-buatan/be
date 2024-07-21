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

	// helper "github.com/cerdas-buatan/be/helper"

	model "github.com/cerdas-buatan/be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// signup
func SignUpPengguna(db *mongo.Database, insertedDoc model.Pengguna) error {
	objectId := primitive.NewObjectID()
	if insertedDoc.Username == "" ||
		insertedDoc.Akun.Password == "" {
		return fmt.Errorf("dimohon untuk melengkapi data")
	}
	if err := checkmail.ValidateFormat(insertedDoc.Akun.Email); err != nil {
		return fmt.Errorf("email tidak valid")
	}
	userExists, _ := GetUserFromEmail(insertedDoc.Akun.Email, db)
	if insertedDoc.Akun.Email == userExists.Email {
		return fmt.Errorf("email sudah terdaftar")
	}
	if strings.Contains(insertedDoc.Akun.Password, " ") {
		return fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Akun.Password) < 8 {
		return fmt.Errorf("password terlalu pendek")
	}
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Akun.Password), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"_id":      objectId,
		"email":    insertedDoc.Akun.Email,
		"password": hex.EncodeToString(hashedPassword),
		"salt":     hex.EncodeToString(salt),
		"role":     "pengguna",
	}
	pengguna := bson.M{
		"username": insertedDoc.Username,
		"akun": model.User{
			ID: objectId,
		},
	}
	_, err = InsertOneDoc(db, "user", user)
	if err != nil {
		return fmt.Errorf("kesalahan server")
	}
	_, err = InsertOneDoc(db, "pengguna", pengguna)
	if err != nil {
		return fmt.Errorf("kesalahan server")
	}
	return nil
}

// login
func LogIn(db *mongo.Database, insertedDoc model.User) (user model.User, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return user, fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return user, fmt.Errorf("email tidak valid")
	}
	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db)
	if err != nil {
		return
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		return user, fmt.Errorf("kesalahan server : salt")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		return user, fmt.Errorf("password salah")
	}
	return existsDoc, nil
}

func GetUserFromEmail(email string, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{"email": email}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("email tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
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

func GetUserFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("no data found for ID %s", _id)
		}
		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return doc, nil
}

// Logout handles user logout (example implementation)
// func Logout(c *fiber.Ctx) error {
// 	// Perform logout logic here, such as clearing session or token
// 	return helper.SendResponse(c, fiber.StatusOK, "Logout successful", nil)
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

func GetAllPengguna(db *mongo.Database) (pengguna []model.Pengguna, err error) {
	collection := db.Collection("pengguna")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return pengguna, fmt.Errorf("error GetAllPengguna mongo: %s", err)
	}
	err = cursor.All(context.Background(), &pengguna)
	if err != nil {
		return pengguna, fmt.Errorf("error GetAllPengguna context: %s", err)
	}
	return pengguna, nil
}

func GetPenggunaFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.Pengguna, err error) {
	collection := db.Collection("pengguna")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("no data found for ID %s", _id)
		}
		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return doc, nil
}

