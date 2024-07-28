package config

import (
	"net/http"
	"os"
)

// Daftar origins yang diizinkan
var Origins = []string{
	"http://cerdas-buatan.projsonal.online/fe",
	"http://cerdas-buatan.projsonal.online/be",
	"http://cerdas-buatan.projsonal.online",
	"http://www.cerdas-buatan.projsonal.online",
	// "https://whatsauth.github.io",
	// "https://www.do.my.id",
}

var Cors = struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
	MaxAge           int
}{
	AllowOrigins:     Origins,
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Content-Type", "Login"},
	AllowCredentials: true,
	ExposeHeaders:    []string{"Content-Length"},
	MaxAge:           3600,
}

// Fungsi untuk memeriksa apakah origin diizinkan
func isAllowedOrigin(origin string) bool {
	for _, o := range Origins {
		if o == origin {
			return true
		}
	}
	return false
}

var Internalhost string = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")

// Fungsi untuk mengatur header CORS
func SetAccessControlHeaders(w http.ResponseWriter, r *http.Request) bool {
	origin := r.Header.Get("Origin")

	if isAllowedOrigin(origin) {
		// Set CORS headers for the preflight request
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Login")
			w.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return true
		}
		// Set CORS headers for the main request.
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		return false
	}
	return false
}

// var Cors= cors.config{
// 	AllowOrigins: Origins,
// 	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 	AllowHeaders: []string{"Content-Type", "Login"},
// 	AllowCredentials: true,
// 	ExposeHeaders: []string{"Content-Length"},
// 	MaxAge: 3600,
// }