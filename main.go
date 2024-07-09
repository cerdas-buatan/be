// package main

// import (
//     "net/http"
//     "github.com/cerdas-buatan/be/module"
// )

// func main() {
//     http.HandleFunc("/chat", handler.ChatHandler)
//     http.ListenAndServe(":8080", nil)
// }
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/cerdas-buatan/be/config"
	"github.com/cerdas-buatan/be/route"
)

func main() {
	app := fiber.New()
	db := config.ConnectDB("your-mongodb-uri")
	route.SetupRoutes(app, db)

	app.Listen(":3000")
}
