package module


func SendWhatsAppVerification(phoneNumber, code string) error {
	// Store the code in Redis
	err := SetVerificationCode(phoneNumber, code)
	if err != nil {
		return err
	}

	// Construct the WhatsApp Web URL to send the message
	message := fmt.Sprintf("Your verification code is: %s", code)
	encodedMessage := url.QueryEscape(message)
	whatsappURL := fmt.Sprintf("https://wa.me/%s?text=%s", phoneNumber, encodedMessage)

	// Open the constructed URL in the user's default browser (simplified approach)
	_, err = http.Get(whatsappURL)
	if err != nil {
		return err
	}
	return nil
}

func ValidateVerificationCode(phoneNumber, code string) bool {
	storedCode, err := GetVerificationCode(phoneNumber)
	if err != nil {
		return false
	}
	return code == storedCode
}
