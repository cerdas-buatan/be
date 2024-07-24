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