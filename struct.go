package be

import (
    "fmt"
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


