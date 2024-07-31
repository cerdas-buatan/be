package helper

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

// Helper function to generate a verification code
func GenerateVerificationCode() string {
	const charset = "1234567890"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Helper function to send a WhatsApp message by automating WhatsApp Web
func SendWhatsAppMessage(phoneNumber, code string) error {
	// Create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run the task
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://web.whatsapp.com"),
		chromedp.Sleep(10*time.Second), // Wait for QR code scanning and WhatsApp Web to load
		chromedp.SendKeys(`//input[@title='Search or start new chat']`, phoneNumber),
		chromedp.Click(`//span[@title='`+phoneNumber+`']`),
		chromedp.SendKeys(`//div[@contenteditable='true']`, "Your verification code is: "+code),
		chromedp.Click(`//span[@data-testid='send']`),
	)
	if err != nil {
		return fmt.Errorf("error sending WhatsApp message: %v", err)
	}

	fmt.Printf("Sent WhatsApp message to %s with code %s\n", phoneNumber, code)
	return nil
}
