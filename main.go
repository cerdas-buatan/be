package be

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/aiteung/atdb"
	"github.com/cerdas-buatan/be/route"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	functions.HTTP("gaysdisal", route.Web)
}

// func main() {
// 	app := fiber.New()
// 	db := config.ConnectDB(os.Getenv("MONGOCONNSTRING"), os.Getenv("DBNAME"))
// 	route.SetupRoutes(app, db)

// 	app.Listen(":3000")
// }

// package main

// import (
//     "net/http"
//     "github.com/cerdas-buatan/be/module"
// )

// func main() {
//     http.HandleFunc("/chat", handler.ChatHandler)
//     http.ListenAndServe(":8080", nil)
// }

func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah"
		}
	}

	return GCFReturnStruct(Response)
}

func InsertUser(db *mongo.Database, collection string, userdata User) string {
	hash, _ := HashPassword(userdata.Password)
	userdata.Password = hash
	atdb.InsertOneDoc(db, collection, userdata)
	return "username : " + userdata.Username + "password : " + userdata.Password
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}
