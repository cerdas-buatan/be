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

// preprocessInput converts input data to tensor
func preprocessInput(input map[string]interface{}) (*tf.Tensor, error) {
	// Customize this function to preprocess input for your model
	// This is a placeholder example
	data := []float32{} // Convert input to the appropriate tensor data
	return tf.NewTensor(data)
}

// postprocessOutput converts model output to the response format
func postprocessOutput(result []*tf.Tensor) (map[string]interface{}, error) {
	// Customize this function to process output from your model
	// This is a placeholder example
	output := result[0].Value().([]float32)
	return map[string]interface{}{
		"prediction": output,
	}, nil
}