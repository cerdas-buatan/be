package model

import (
    "fmt"
    "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Register struct {
    Email           string
    Username        string
    Password        string
    ConfirmPassword string
}

type Login struct {
    Username    string
    Password    string
}


