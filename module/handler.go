package module

import (
	"encoding/json"
	"net/http"
	"os"

	model "github.com/cerdas-buatan/be/model"
	"github.com/cerdas-buatan/be/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var (
// 	Response model.Response
// 	user model.User
// 	pengguna model.Pengguna
// 	driver model.Driver
// 	obat model.Obat
// 	order model.Order
// 	password model.Password

// )

var (
	Response model.Response
	// user     model.User
	// pengguna model.Pengguna
	// password model.Password
)

// signup
func GCFHandlerSignUpPengguna(MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.Response
	Response.Status = false
	var datapengguna model.Pengguna
	err := json.NewDecoder(r.Body).Decode(&datapengguna)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(Response)
	}
	err = SignUpPengguna(conn, datapengguna)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	Response.Message = "Halo " + datapengguna.NamaLengkap
	return GCFReturnStruct(Response)
}

// login
func GCFHandlerLogin(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.Credential
	Response.Status = false
	var datauser model.User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(Response)
	}
	user, err := LogIn(conn, datauser)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	tokenstring, err := Encode(user.ID, user.Role, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		Response.Message = "Gagal Encode Token : " + err.Error()
	} else {
		Response.Message = "Selamat Datang " + user.Email
		Response.Token = tokenstring
		Response.Role = user.Role
	}
	return GCFReturnStruct(Response)
}

// get all
func GCFHandlerGetAll(MONGOCONNSTRINGENV, dbname, col string, docs interface{}) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	data := GetAllDocs(conn, col, docs)
	return GCFReturnStruct(data)
}



func GCFHandlerGetUserFromID(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.Response
	Response.Status = false
	tokenstring := r.Header.Get("Authorization")
	payload, err := Decode(os.Getenv(PASETOPUBLICKEYENV), tokenstring)
	if err != nil {
		Response.Message = "Gagal Decode Token : " + err.Error()
		return GCFReturnStruct(Response)
	}
	data, err := GetUserFromID(payload.Id, conn)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}
	return GCFReturnStruct(data)
}


func GCFHandlerUpdateByPengguna(idparam, iduser primitive.ObjectID, pengguna model.Pengguna, conn *mongo.Database, r *http.Request) string {
	Response.Status = false
	//
	err := UpdatePengguna(idparam, iduser, conn, pengguna)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}
	//
	Response.Status = true
	Response.Message = "Berhasil Update Pengguna"
	return GCFReturnStruct(Response)
}


func GCFHandlerGetAllPengguna(MONGOCONNSTRINGENV, dbname string) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	var Response model.Response
	Response.Status = false
	data, err := GetAllPengguna(conn)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}
	return GCFReturnStruct(data)
}



func GCFHandlerGetPenggunaByPengguna(iduser primitive.ObjectID, conn *mongo.Database) string {
	Response.Status = false
	//
	pengguna, err := GetPenggunaFromAkun(iduser, conn)
	if err != nil {
		Response.Message = err.Error()
		return GCFReturnStruct(Response)
	}
	//
	return GCFReturnStruct(pengguna)
}

//qna
func GCFHandlerQnA(r *http.Request) string {
	var Response model.Response
	Response.Status = false
	var qnaRequest struct {
		Question string `json:"question"`
	}
	err := json.NewDecoder(r.Body).Decode(&qnaRequest)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(Response)
	}

	answer, err := model.GetAnswer(qnaRequest.Question)
	if err != nil {
		Response.Message = "error getting answer: " + err.Error()
		return GCFReturnStruct(Response)
	}

	Response.Status = true
	Response.Message = answer
	return GCFReturnStruct(Response)
}



// return struct
func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

// get user login
func GetUserLogin(PASETOPUBLICKEYENV string, r *http.Request) (model.Payload, error) {
	tokenstring := r.Header.Get("Authorization")
	payload, err := Decode(os.Getenv(PASETOPUBLICKEYENV), tokenstring)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

// get id
func GetID(r *http.Request) string {
	return r.URL.Query().Get("id")
}
