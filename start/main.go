// initialize the database and routes for local development.
package main

import (
	"log"
	"net/http"

	"github.com/cerdas-buatan/be/config"
	"github.com/cerdas-buatan/be/module"
	"github.com/cerdas-buatan/be/route"
)

func main() {
	// Initialize database connection and services
	db := config.InitDB()
	menuService := module.NewMenuService(db)

	// Initialize routes
	url.InitRoutes(menuService)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
