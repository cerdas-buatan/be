package helper

import "encoding/json"

func WriteJSON(respw http.ResponseWriter, statusCode int, content interface{}) {
    respw.Header().Set("Content-Type", "application/json")
    respw.WriteHeader(statusCode)
    respw.Write([]byte(Jsonstr(content)))
}

func Jsonstr(strc interface{}) string {
    jsonData, err := json.Marshal(strc)
    if (err != nil) {
        log.Fatal(err)
    }
    return string(jsonData)
}




// func GCFReturnStruct(DataStruct interface{}) string {
// 	jsonData, _ := json.Marshal(DataStruct)
// 	return string(jsonData)
// }











// package main
// import(
// 	"net/http"
// 	"github.com/gofiber/fiber/v2/log"
// 	"encoding/json"
// )

// func WriteJSON(respw http.ResponseWriter, statusCode int, content interface{}) {
//     respw.Header().Set("Content-Type", "application/json")
//     respw.WriteHeader(statusCode)
//     respw.Write([]byte(Jsonstr(content)))
// }

// func Jsonstr(strc interface{}) string {
//     jsonData, err := json.Marshal(strc)
//     if (err != nil) {
//         log.Fatal(err)
//     }
//     return string(jsonData)
// }

// func SendResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
// 	return c.Status(statusCode).JSON(fiber.Map{
// 		"message": message,
// 		"data":    data,
// 	})
// }
