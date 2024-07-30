package helper

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"strings"

	model "github.com/cerdas-buatan/be/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/argon2"
)

// hash
func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return hex.EncodeToString(hashedPassword), nil
}

// sendresponse
func SendResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	response := model.Response{
		Status:  status == fiber.StatusOK,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(response)
}

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

// KOMEN FUNCT DIBAWAH BOLEH DIHAPUS 

// GetEnv gets the value of an environment variable or returns a default value if not set
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