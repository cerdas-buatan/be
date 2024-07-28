package helper

import (
	"fmt"
	"math/rand"

	"github.com/whatsauth/whatsauth"
)

// Helper function to generate a verification code
func generateVerificationCode() string {
	const charset = "1234567890"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Helper function to send a WhatsApp message using the whatsauth library
func sendWhatsAppMessage(phoneNumber, code string) error {
	// Use the whatsauth library to send the message
	waToken := whatsauth.New(phoneNumber, code)
	err := waToken.Send("Your WhatsApp API Key")
	if err != nil {
		return fmt.Errorf("error sending WhatsApp message: %v", err)
	}
	fmt.Printf("Sending WhatsApp message to %s with code %s\n", phoneNumber, code)
	return nil
}
