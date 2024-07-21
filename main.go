package be

import (
	"encoding/json"
	"net/http"


	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	route "github.com/cerdas-buatan/be/route"
)

func init() {
	functions.HTTP("gaysdisal", route.Web)
}

func HomeGaysdisal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Welcome to Gaysdisal",
	})
}

