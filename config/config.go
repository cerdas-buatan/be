package config

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"os"
)

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Gaysdisal",
	AppName:       "Gaysdisal",
}

func GetAddress() (ipport string, network string) {
	port := os.Getenv("PORT")
	ip := os.Getenv("IP")

	// Default values
	if port == "" {
		port = ":8080"
	} else if port[0:1] != ":" {
		port = ":" + port
	}

	network = "tcp4"
	ipport = port

	if ip != "" {
		if strings.Contains(ip, ".") {
			ipport = ip + port
		} else {
			ipport = "[" + ip + "]" + port
			network = "tcp6"
		}
	}

	return ipport, network
}

var IPort, netString = GetAddress()

var PrivateKey = os.Getenv("PRIVATEKEY")
var PublicKey = os.Getenv("PUBLICKEY")