// package module


// func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	datagedung := GetAllBangunanLineString(mconn, collectionname)
// 	return GCFReturnStruct(datagedung)
// }

// func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
// 	var Response Credential
// 	Response.Status = false
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	var datauser User
// 	err := json.NewDecoder(r.Body).Decode(&datauser)
// 	if err != nil {
// 		Response.Message = "error parsing application/json: " + err.Error()
// 	} else {
// 		if IsPasswordValid(mconn, collectionname, datauser) {
// 			Response.Status = true
// 			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
// 			if err != nil {
// 				Response.Message = "Gagal Encode Token : " + err.Error()
// 			} else {
// 				Response.Message = "Selamat Datang"
// 				Response.Token = tokenstring
// 			}
// 		} else {
// 			Response.Message = "Password Salah"
// 		}
// 	}

// 	return GCFReturnStruct(Response)
// }

// func GCFReturnStruct(DataStuct any) string {
// 	jsondata, _ := json.Marshal(DataStuct)
// 	return string(jsondata)
// }

// func InsertUser(db *mongo.Database, collection string, userdata User) string {
// 	hash, _ := HashPassword(userdata.Password)
// 	userdata.Password = hash
// 	atdb.InsertOneDoc(db, collection, userdata)
// 	return "Ini username : " + userdata.Username + "ini password : " + userdata.Password
// }

package module

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"net/http"
	"encoding/json"

	"github.com/cerdas-buatan/be/model"
	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/argon2"
)

// var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(MongoString, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// crud
func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server : insert")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}

func DeleteOneDoc(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// register
// Handler untuk menangani registrasi pengguna
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    // Mendapatkan data registrasi dari body request
    var registerData model.Register
    err := json.NewDecoder(r.Body).Decode(&registerData)
    if err != nil {
        http.Error(w, "Gagal memproses data registrasi", http.StatusBadRequest)
        return
    }

    // Validasi data registrasi
    if registerData.Email == "" || registerData.Username == "" ||
        registerData.Password == "" || registerData.ConfirmPassword == "" {
        http.Error(w, "Dimohon untuk melengkapi data", http.StatusBadRequest)
        return
    }
    if registerData.Password != registerData.ConfirmPassword {
        http.Error(w, "Konfirmasi password tidak cocok", http.StatusBadRequest)
        return
    }

    // Validasi format email menggunakan package checkmail
    if err := checkmail.ValidateFormat(registerData.Email); err != nil {
        http.Error(w, "Email tidak valid", http.StatusBadRequest)
        return
    }

    // Validasi apakah email sudah terdaftar
    userExists, err := GetUserFromEmail(registerData.Email, database)
    if err == nil && userExists.Email == registerData.Email {
        http.Error(w, "Email sudah terdaftar", http.StatusBadRequest)
        return
    }

    // Validasi password tidak boleh mengandung spasi
    if strings.Contains(registerData.Password, " ") {
        http.Error(w, "Password tidak boleh mengandung spasi", http.StatusBadRequest)
        return
    }

    // Validasi panjang minimal password
    if len(registerData.Password) < 8 {
        http.Error(w, "Password terlalu pendek", http.StatusBadRequest)
        return
    }

    // Generate ObjectId untuk pengguna baru
    userID := primitive.NewObjectID()

    // Hash password sebelum menyimpannya ke database
    salt := make([]byte, 16)
    _, err = rand.Read(salt)
    if err != nil {
        http.Error(w, "Gagal menghasilkan salt", http.StatusInternalServerError)
        return
    }
    hashedPassword := argon2.IDKey([]byte(registerData.Password), salt, 1, 64*1024, 4, 32)

    // Persiapkan data pengguna untuk disimpan di MongoDB
    user := bson.M{
        "_id":      userID,
        "email":    registerData.Email,
        "username": registerData.Username,
        "password": hex.EncodeToString(hashedPassword),
        "salt":     hex.EncodeToString(salt),
        "role":     "pengguna",
    }

    // Simpan data pengguna ke koleksi 'users' di MongoDB
    _, err = InsertOneDoc(database, "users", user)
    if err != nil {
        http.Error(w, "Gagal menyimpan data pengguna", http.StatusInternalServerError)
        return
    }

    // Kirim respons berhasil
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Registrasi berhasil"})
}

//login
func LogIn(db *mongo.Database, insertedDoc model.User) (user model.User, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return user, fmt.Errorf("Dimohon untuk melengkapi data")
	} 
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return user, fmt.Errorf("Email tidak valid")
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

//user
func UpdateEmailUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.User) error {
	dataUser, err := GetUserFromID(iduser, db)
	if err != nil {
		return err
	}
	if insertedDoc.Email == "" {
		return fmt.Errorf("Dimohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return fmt.Errorf("Email tidak valid")
	}
	existsDoc, _ := GetUserFromEmail(insertedDoc.Email, db)
	if existsDoc.Email == insertedDoc.Email {
		return fmt.Errorf("Email sudah terdaftar")
	}
	user := bson.M{
		"email": insertedDoc.Email,
		"password": dataUser.Password,
		"salt": dataUser.Salt,
		"role": dataUser.Role,
	}
	err = UpdateOneDoc(iduser, db, "user", user)
	if err != nil {
		return err
	}
	return nil
}

// func UpdateUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.User) error {
// 	dataUser, err := GetUserFromID(iduser, db)
// 	if err != nil {
// 		return err
// 	}
// 	if insertedDoc.Email == "" || insertedDoc.Password == "" {
// 		return fmt.Errorf("Dimohon untuk melengkapi data")
// 	}
// 	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
// 		return fmt.Errorf("Email tidak valid")
// 	}
// 	existsDoc, _ := GetUserFromEmail(insertedDoc.Email, db)
// 	if existsDoc.Email == insertedDoc.Email {
// 		return fmt.Errorf("Email sudah terdaftar")
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
// 		"email": insertedDoc.Email,
// 		"password": hex.EncodeToString(hashedPassword),
// 		"salt": hex.EncodeToString(salt),
// 		"role": dataUser.Role,
// 	}
// 	err = UpdateOneDoc(iduser, db, "user", user)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func UpdatePasswordUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.Password) error {
	dataUser, err := GetUserFromID(iduser, db)
	if err != nil {
		return err
	}
	salt, err := hex.DecodeString(dataUser.Salt)
	if err != nil {
		return fmt.Errorf("kesalahan server : salt")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != dataUser.Password {
		return fmt.Errorf("password lama salah")
	}
	if insertedDoc.Newpassword == ""  {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	if strings.Contains(insertedDoc.Newpassword, " ") {
		return fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Newpassword) < 8 {
		return fmt.Errorf("password terlalu pendek")
	}
	salt = make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Newpassword), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"email": dataUser.Email,
		"password": hex.EncodeToString(hashedPassword),
		"salt": hex.EncodeToString(salt),
		"role": dataUser.Role,
	}
	err = UpdateOneDoc(iduser, db, "user", user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.User) error {
	dataUser, err := GetUserFromID(iduser, db)
	if err != nil {
		return err
	}
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return fmt.Errorf("email tidak valid")
	}
	existsDoc, _ := GetUserFromEmail(insertedDoc.Email, db)
	if existsDoc.Email == insertedDoc.Email {
		return fmt.Errorf("email sudah terdaftar")
	}
	if strings.Contains(insertedDoc.Password, " ") {
		return fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Password) < 8 {
		return fmt.Errorf("password terlalu pendek")
	}
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"email": insertedDoc.Email,
		"password": hex.EncodeToString(hashedPassword),
		"salt": hex.EncodeToString(salt),
		"role": dataUser.Role,
	}
	err = UpdateOneDoc(iduser, db, "user", user)
	if err != nil {
		return err
	}
	return nil
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

// pengguna
func UpdatePengguna(idparam, iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.Pengguna) error {
	pengguna, err := GetPenggunaFromAkun(iduser, db)
	if err != nil {
		return err
	}
	if pengguna.ID != idparam {
		return fmt.Errorf("Anda bukan pemilik data ini")
	}
	if insertedDoc.NamaLengkap == "" || insertedDoc.TanggalLahir == "" || insertedDoc.JenisKelamin == "" || insertedDoc.NomorHP == "" || insertedDoc.Alamat == ""{
		return fmt.Errorf("Dimohon untuk melengkapi data")
	} 
	pgn := bson.M{
		"namalengkap": insertedDoc.NamaLengkap,
		"tanggallahir": insertedDoc.TanggalLahir,
		"jeniskelamin": insertedDoc.JenisKelamin,
		"nomorhp": insertedDoc.NomorHP,
		"alamat": insertedDoc.Alamat,
		"akun": model.User {
			ID : pengguna.Akun.ID,
		},
	}
	err = UpdateOneDoc(idparam, db, "pengguna", pgn)
	if err != nil {
		return err
	}
	return nil
}

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

func GetPenggunaFromAkun(akun primitive.ObjectID, db *mongo.Database) (doc model.Pengguna, err error) {
	collection := db.Collection("pengguna")
	filter := bson.M{"akun._id": akun}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("pengguna tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	return doc, nil
}

//by admin
func GetPenggunaFromIDByAdmin(idparam primitive.ObjectID, db *mongo.Database) (pengguna model.Pengguna, err error) {
	collection := db.Collection("pengguna")
	filter := bson.M{
		"_id": idparam,
	}
	err = collection.FindOne(context.Background(), filter).Decode(&pengguna)
	if err != nil {
		return pengguna, fmt.Errorf("error GetPenggunaFromID mongo: %s", err)
	}
	user, err := GetUserFromID(pengguna.Akun.ID, db)
	if err != nil {
		return pengguna, fmt.Errorf("error GetPenggunaFromID mongo: %s", err)
	}
	akun := model.User{
		ID: user.ID,
		Email: user.Email,
		Role: user.Role,
	}
	pengguna.Akun = akun
	return pengguna, nil
}

func GetAllPenggunaByAdmin(db *mongo.Database) (pengguna []model.Pengguna, err error) {
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

// verifikasi
func VerifyAfterLoginHandler(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan userID dari sesi atau request, asumsikan sudah ada
	userID := primitive.NewObjectID() // Ganti dengan userID sesuai implementasi Anda

	// Panggil fungsi verifikasi
	err := VerifyAfterLogin(database, userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("gagal verifikasi: %v", err), http.StatusInternalServerError)
		return
	}

	// Kirim respons berhasil
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Verifikasi berhasil disimpan"})
}

