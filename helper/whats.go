package helper
import{
	"math/rand"
	"fmt"
}

// Helper function to generate a verification code
func generateVerificationCode() string {
	const charset = "1234567890"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Helper function to send a WhatsApp message (add your own implementation)
func sendWhatsAppMessage(phoneNumber, code string) error {
	// Implement your WhatsApp API integration here
	// Example: Call WhatsApp API with the phoneNumber and code
	fmt.Printf("Sending WhatsApp message to %s with code %s\n", phoneNumber, code)
	return nil
}