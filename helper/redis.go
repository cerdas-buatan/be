package helper

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var rdb *redis.Client

func InitRedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

func SetVerificationCode(phoneNumber, code string) error {
	ctx := context.Background()
	return rdb.Set(ctx, phoneNumber, code, 5*time.Minute).Err() // Code expires in 5 minutes
}

func GetVerificationCode(phoneNumber string) (string, error) {
	ctx := context.Background()
	return rdb.Get(ctx, phoneNumber).Result()
}
