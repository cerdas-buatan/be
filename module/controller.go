package module

import (
	"context"
	"encoding/hex"
//	"context"
//	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

//	"github.com/badoux/checkmail"
//	"golang.org/x/crypto/argon2"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/argon2"

	model "github.com/cerdas-buatan/be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Helper function to generate BoW representation
func generateBoW(message string) map[string]int {
	words := strings.Fields(message)
	bow := make(map[string]int)
	for _, word := range words {
		bow[word]++
	}
	return bow
}







// Function to send BoW to the IndoBERT API for prediction
func getPredictionFromIndoBERT(bow map[string]int) (string, error) {
	// Convert BoW to JSON
	jsonData, err := json.Marshal(bow)
	if err != nil {
		return "", fmt.Errorf("error marshalling BoW to JSON: %v", err)
	}


	// Send request to the IndoBERT API
	resp, err := http.Post("YOUR_INDOBERT_API_URL", "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return "", fmt.Errorf("error sending request to IndoBERT API: %v", err)
	}
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("IndoBERT API returned status: %v", resp.StatusCode)
	}


	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding IndoBERT API response: %v", err)
	}


	return result["response"], nil
}


// GCFGetResponse generates a response based on the input message using BoW and IndoBERT
func GCFGetResponse(message string, db *mongo.Database) (string, error) {
	// Generate BoW representation
	bow := generateBoW(message)


	// Get prediction from IndoBERT API
	response, err := getPredictionFromIndoBERT(bow)
	if err != nil {
		return "", err
	}


	return response, nil
}



// ChatHandler handles chat requests and generates responses using IndoBERT
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var chatReq model.ChatRequest
	err := json.NewDecoder(r.Body).Decode(&chatReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := GCFGetResponse(chatReq.Message, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chatRes := model.ChatResponse{Response: response}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chatRes)
}

