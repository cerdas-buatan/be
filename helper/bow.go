package helper
import (
	"math"
	"strings"
)

// Tokenize splits a string into words
func tokenize(s string) []string {
	return strings.Fields(s)
}

// CreateVocab creates a vocabulary of unique words from two strings
func createVocab(s1, s2 string) map[string]bool {
	vocab := make(map[string]bool)
	for _, word := range tokenize(s1) {
		vocab[word] = true
	}
	for _, word := range tokenize(s2) {
		vocab[word] = true
	}
	return vocab
}

// Vectorize creates a word count vector from a string and a vocabulary
func vectorize(s string, vocab map[string]bool) map[string]int {
	vector := make(map[string]int)
	words := tokenize(s)
	for word := range vocab {
		vector[word] = 0
	}
	for _, word := range words {
		vector[word]++
	}
	return vector
}

// CosineSimilarity computes the cosine similarity between two vectors
func cosineSimilarity(v1, v2 map[string]int) float64 {
	var dotProduct, mag1, mag2 float64
	for word, count1 := range v1 {
		count2 := v2[word]
		dotProduct += float64(count1 * count2)
		mag1 += float64(count1 * count1)
		mag2 += float64(count2 * count2)
	}
	if mag1 == 0 || mag2 == 0 {
		return 0.0
	}
	return dotProduct / (math.Sqrt(mag1) * math.Sqrt(mag2))
}

// 
// func BagOfWordsSimilarity(s1, s2 string) float64 {
// 	vocab := createVocab(s1, s2)
// 	vector1 := vectorize(s1, vocab)
// 	vector2 := vectorize(s2, vocab)
// 	return cosineSimilarity(vector1, vector2)
// }

func BagOfWordsSimilarity(s1, s2 string) float64 {
		vocab := createVocab(s1, s2)
		vector1 := vectorize(s1, vocab)
		vector2 := vectorize(s2, vocab)
		return cosineSimilarity(vector1, vector2)
	}
