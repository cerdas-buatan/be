package module
import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/tensorflow/tensorflow/tensorflow/go"
)

// predictHandler handles prediction requests
func predictHandler(w http.ResponseWriter, r *http.Request) {
	// Load the TensorFlow model
	model, err := tf.LoadSavedModel("indobert_model", []string{"serve"}, nil)
	if err != nil {
		http.Error(w, "Failed to load model", http.StatusInternalServerError)
		return
	}
	defer model.Session.Close()

	// Decode input data from request
	var input map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Preprocess input
	// Note: Customize preprocessing according to your model's needs
	// Here is an example of converting input to tensor
	tensor, err := preprocessInput(input)
	if err != nil {
		http.Error(w, "Failed to preprocess input", http.StatusInternalServerError)
		return
	}

	// Run the model
	result, err := model.Session.Run(
		[]*tf.Tensor{tensor},
		[]string{"output_node_name"}, // Replace with your model's output node
		nil,
	)
	if err != nil {
		http.Error(w, "Failed to run model", http.StatusInternalServerError)
		return
	}

	// Postprocess output
	response, err := postprocessOutput(result)
	if err != nil {
		http.Error(w, "Failed to postprocess output", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// file ini mau dihapus silakan