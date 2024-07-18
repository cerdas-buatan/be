package config

import (
	"github.com/gofiber/fiber/v2"
	// helper "github.com/cerdas-buaatan/be/helper"
	"os"
)

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Gaysdisal",
	AppName:       "Gaysdisal",
}
var IPort, netString = helper.GetAddress()

var PrivateKey = os.Getenv("PRIVATEKEY")
var PublicKey = os.Getenv("PUBLICKEY")