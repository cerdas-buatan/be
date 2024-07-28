package helper

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/whatsauth/whatsauth"
	"os"
	"time"
)

var rdb *redis.Client
var waClient *whatsauth.Client

func InitRedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

func InitWhatsAuthClient() {
	waClient = whatsauth.NewClient(os.Getenv("WHATSAPP_API_KEY"))
}

func SetVerificationCode(phoneNumber, code string) error {
	ctx := context.Background()
	return rdb.Set(ctx, phoneNumber, code, 5*time.Minute).Err() // Code expires in 5 minutes
}

func GetVerificationCode(phoneNumber string) (string, error) {
	ctx := context.Background()
	return rdb.Get(ctx, phoneNumber).Result()
}

func SendWhatsAppMessage(phoneNumber, code string) error {
	message := fmt.Sprintf("Your verification code is %s", code)
	waToken, err := waClient.Send(phoneNumber, message)
	if err != nil {
		return fmt.Errorf("error sending WhatsApp message: %v", err)
	}
	fmt.Printf("Sent WhatsApp message to %s with code %s\n", phoneNumber, code)
	return waToken.Send()
}
