package helper

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

// NewMessage creates a new WhatsApp text message
func NewMessage(phoneNumber, code string) whatsapp.TextMessage {
	return whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: phoneNumber + "@s.whatsapp.net",
		},
		Text: "Your verification code is: " + code,
	}
}

// sendMessage sends a WhatsApp message using the go-whatsapp library
func sendMessage(msg whatsapp.TextMessage) error {
	// Create a new WhatsApp connection
	wac, err := whatsapp.NewConn(20 * time.Second)
	if err != nil {
		return fmt.Errorf("error creating WhatsApp connection: %v", err)
	}

	// Handle the QR code for login
	qrChan := make(chan string)
	go func() {
		// Display the QR code
		for qr := range qrChan {
			fmt.Printf("Scan the following QR code to login: %v\n", qr)
		}
	}()
	session, err := wac.Login(qrChan)
	if err != nil {
		return fmt.Errorf("error logging into WhatsApp: %v", err)
	}
	log.Printf("Logged into WhatsApp as %v\n", session.Wid)

	// Send the message
	_, err = wac.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending WhatsApp message: %v", err)
	}

	fmt.Printf("Sent WhatsApp message to %s\n", msg.Info.RemoteJid)
	return nil
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
