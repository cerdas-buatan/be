package module

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aiteung/atdb"
	model "github.com/cerdas-buatan/be/model"
	"github.com/whatsauth/watoken"

	// helper"github.com/cerdas-buatan/be/helper"
	// "github.com/aiteung/atdb"

	"go.mongodb.org/mongo-driver/mongo"
)

// GCFHandlerSignUpPengguna handles signup for Google Cloud Function
func GCFSignUp(w http.ResponseWriter, r *http.Request) {
	db := helper.ConnectMongoDB(os.Getenv("MONGOCONNSTRING"), os.Getenv("DBNAME"))
	defer db.Client().Disconnect(r.Context())

	var Response model.Response
	Response.Status = false

	var userdata model.Pengguna
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response)
		return
	}

	err = helper.SignUpPengguna(db, userdata)
	if err != nil {
		Response.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response)
		return
	}

	Response.Status = true
	Response.Message = "Halo, Selamat Bergabung " + userdata.Username
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response)
}

// GCFHandlerLogin handles login for Google Cloud Function
func GCFSignIn(w http.ResponseWriter, r *http.Request) {
	db := helper.ConnectMongoDB(os.Getenv("MONGOCONNSTRING"), os.Getenv("DBNAME"))
	defer db.Client().Disconnect(r.Context())

	var Response model.Credential
	Response.Status = false

	var datapengguna model.User
	err := json.NewDecoder(r.Body).Decode(&datapengguna)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response)
		return
	}

	user, err := helper.LogIn(db, datapengguna)
	if err != nil {
		Response.Message = err.Error()
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response)
		return
	}

	Response.Status = true
	tokenstring, err := helper.Encode(user.ID, user.Role, os.Getenv("PASETOPRIVATEKEYENV"))
	if err != nil {
		Response.Message = "Gagal Encode Token : " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		Response.Message = "Halo, Selamat Datang " + user.Email
		Response.Token = tokenstring
		Response.Role = user.Role
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response)
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

// ChatHandler handles chat requests
func GCFChat(MongoString, dbname string, w http.ResponseWriter, r *http.Request) {
	var Reqchat model.ChatRequest
	err := json.NewDecoder(r.Body).Decode(&Reqchat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := helper.MongoConnect(MongoString, dbname)
	defer func() {
		if err := db.Client().Disconnect(context.TODO()); err != nil {
			log.Printf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	response, err := GetResponse(Reqchat.Message, db) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Reqchat := model.ChatResponse{Response: response}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Reqchat)
}

func InsertUser(db *mongo.Database, collection string, userdata User) string {
	hash, _ := HashPassword(userdata.Password)
	userdata.Password = hash
	atdb.InsertOneDoc(db, collection, userdata)
	return "username : " + userdata.Username + "password : " + userdata.Password
}

// get all
func GCFHandlerGetAll(MONGOCONNSTRINGENV, dbname, col string, docs interface{}) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	data := GetAllDocs(conn, col, docs)
	return GCFReturnStruct(data)
}

func GCFPredict(w http.ResponseWriter, r *http.Request) {
	predictHandler(w, r)
}

func GCFGetUserFromID(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
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

