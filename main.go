package be

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cerdas-buatan/be/route"
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
