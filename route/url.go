package url

import (
	"github.com/cerdas-buatan/be/config"
	"github.com/cerdas-buatan/be/module"
	"net/http"
)

func Web(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return // If it's a preflight request, return early.
	}
	var method, path string = r.Method, r.URL.Path
	switch {
	case method == "GET" && path == "/":
		controller.HomeMakmur(w, r)
	case method == "POST" && path == "/registerai":
		controller.RegisterUsers(w, r)
	case method == "POST" && path == "/loginai":
		controller.LoginUsers(w, r)
	case method == "GET" && path == "/getuser":
		controller.GetUser(w, r)
	case method == "GET" && path == "/getallusers":
		controller.GetAllUsers(w, r)
	case method == "POST" && path == "/chatGaysdisal":
		controller.ChatPredict(w, r)
	case method == "POST" && path == "/chatRegex":
		controller.ChatPredictRegex(w, r)
	default:
		controller.NotFound(w, r)
	}
}

func InitRoutes(menuService *module.MenuService) {
	MenuRoutes(menuService) // Initialize menu routes
	http.HandleFunc("/", Web) // Initialize main routes
}
