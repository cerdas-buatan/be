package route

import (
	"net/http"

	module "github.com/cerdas-buatan/be/module"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func SetupRouter(app *fiber.App, database *mongo.Database) {
	db = database

	app.Get("/", func(c *fiber.Ctx) error {
		return module.HomeGaysdisal(c, db)
	})

	app.Post("/registerai", func(c *fiber.Ctx) error {
		return module.RegisterUsers(c, db)
	})

	app.Post("/loginai", func(c *fiber.Ctx) error {
		return module.LoginUsers(c, db)
	})

	app.Use(func(c *fiber.Ctx) error {
		return module.NotFound(c)
	})
}

func Web(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()
	SetupRouter(app, db)
	app.Handler()(w, r)
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
