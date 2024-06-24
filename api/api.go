package api

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type PredictionResponse struct {
	Prediction int `json:"prediction"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/predict", func(w http.ResponseWriter, r *http.Request) {
		question := r.URL.Query().Get("question")
		if question == "" {
			http.Error(w, "Missing question parameter", http.StatusBadRequest)
			return
		}

		cmd := exec.Command("python3", "predict.py", question)
		output, err := cmd.Output()
		if err != nil {
			http.Error(w, "Failed to execute model", http.StatusInternalServerError)
			return
		}

		var prediction PredictionResponse
		if err := json.Unmarshal(output, &prediction); err != nil {
			http.Error(w, "Failed to parse model output", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(prediction)
	})

	http.ListenAndServe(":8080", r)
}
