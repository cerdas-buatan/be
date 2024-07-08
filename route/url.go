package route

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
		module.HomeMakmur(w, r)
	case method == "POST" && path == "/registerai":
		module.RegisterUsers(w, r)
	case method == "POST" && path == "/loginai":
		module.LoginUsers(w, r)
	default:
		module.NotFound(w, r)
	}
}