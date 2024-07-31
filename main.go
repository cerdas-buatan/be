// main.go
package main

import (
    "net/http"
    "log"
    "github.com/cerdas-buatan/be/config"
    "github.com/cerdas-buatan/be/module"
    "github.com/cerdas-buatan/be/route"
	"github.com/cerdas-buatan/be/model"
    "github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
    functions.HTTP("gaysdisal", route.Web)
}

func main() {
    db := config.InitDB()
    menuService := module.NewMenuService(db)
    route.InitRoutes(menuService)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
