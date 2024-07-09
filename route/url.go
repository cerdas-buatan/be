// package route

// import (
// 	"github.com/cerdas-buatan/be/config"
// 	"github.com/cerdas-buatan/be/module"
// 	"net/http"
// 	"github.com/gofiber/fiber/v2"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func Web(w http.ResponseWriter, r *http.Request) {
// 	if config.SetAccessControlHeaders(w, r) {
// 		return // If it's a preflight request, return early.
// 	}
// 	var method, path string = r.Method, r.URL.Path
// 	switch {
// 	case method == "GET" && path == "/":
// 		module.HomeMakmur(w, r)
// 	case method == "POST" && path == "/registerai":
// 		module.RegisterUsers(w, r)
// 	case method == "POST" && path == "/loginai":
// 		module.LoginUsers(w, r)
// 	default:
// 		module.NotFound(w, r)
// 	}
// }

package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/cerdas-buatan/be/config"
	"github.com/cerdas-buatan/be/module"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, db *mongo.Database) {
	config.SetupCors(app)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Post("/register", module.RegisterUser)
	app.Post("/signup", module.GCFHandlerSignUpPengguna)
	app.Post("/login", module.GCFHandlerLogin)
	app.Post("/chat", module.ChatHandler)
}
