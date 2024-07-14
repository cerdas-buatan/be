package route

import (
	"github.com/gofiber/fiber/v2"
	config "github.com/cerdas-buatan/be/config"
	module "github.com/cerdas-buatan/be/module"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, db *mongo.Database) {
	config.SetupCors(app)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// app.Post("/registeruser", module.RegisterUser)
	app.Post("/registerai", module.GCFHandlerSignUpPengguna)
	app.Post("/loginai", module.GCFHandlerLogin)
	app.Post("/chatres", module.ChatHandler)
}






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
// 		module.HomeGaysdisal(w, r)
// 	case method == "POST" && path == "/registerai":
// 		module.RegisterUsers(w, r)
// 	case method == "POST" && path == "/loginai":
// 		module.LoginUsers(w, r)
// 	default:
// 		module.NotFound(w, r)
// 	}
// }

