package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func Jsonstr(strc interface{}) string {
	jsonData, err := json.Marshal(strc)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}

func WriteJSON(respw http.ResponseWriter, statusCode int, content interface{}) {
	respw.Header().Set("Content-Type", "application/json")
	respw.WriteHeader(statusCode)
	respw.Write([]byte(Jsonstr(content)))
}