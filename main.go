package main


import (
	"encoding/json"
	"net/http"

	"log"
	"net/http"

	"github.com/cerdas-buatan/be/config"
	"github.com/cerdas-buatan/be/module"
	"github.com/cerdas-buatan/be/route"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	route "github.com/cerdas-buatan/be/route"
)


func init() {
	functions.HTTP("gaysdisal", route.Web)
}

func main() {
	// Initialize database connection and services
	db := config.InitDB()
	menuService := module.NewMenuService(db)

	// Initialize routes
	url.InitRoutes(menuService)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}