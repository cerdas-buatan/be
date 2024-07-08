package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GetAddress is a function to get IP and Port from environment variable
func GetAddress() (ipport string, network string) {
	port := os.Getenv("PORT")
	ip := os.Getenv("IP")

	// Default values
	if port == "" {
		port = ":8080"
	} else if port[0:1] != ":" {
		port = ":" + port
	}

	network = "tcp4"
	ipport = port

	if ip != "" {
		if strings.Contains(ip, ".") {
			ipport = ip + port
		} else {
			ipport = "[" + ip + "]" + port
			network = "tcp6"
		}
	}

	return ipport, network
}

// // GetEnv gets the value of an environment variable or returns a default value if not set
// func GetEnv(key, defaultValue string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}
// 	return defaultValue
// }

// // IsIPv6 checks if the given IP address is IPv6
// func IsIPv6(ip string) bool {
// 	return strings.Contains(ip, ":")
// }

// // GenerateUUID generates a new UUID
// func GenerateUUID() string {
// 	return uuid.New().String()
// }

// // HashPassword hashes the given password using bcrypt
// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }

// // CheckPasswordHash checks if the given password matches the hashed password
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// // HashString hashes a string using SHA-256
// func HashString(input string) string {
// 	hash := sha256.New()
// 	hash.Write([]byte(input))
// 	return hex.EncodeToString(hash.Sum(nil))
// }
