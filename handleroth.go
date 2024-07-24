package module

import (
	"encoding/json"
	"net/http"
	"os"
//	"encoding/json"
//	"net/http"
//	"os"
	"github.com/aiteung/atdb"
	model "github.com/cerdas-buatan/be/model"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/mongo"
)

// GCFHandlerSignUpPengguna handles signup for Google Cloud Function
func GCFHandlerSignUp(MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.Response
	Response.Status = false
	var datauser model.Pengguna
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(Response)
	}
	err = SignUp(conn, datauser)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Halo, Selamat Bergabung" + datauser.Username
	return GCFReturnStruct(Response)
}

// signup
func SignUp(db *mongo.Database, insertedDoc model.Pengguna) error {
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
