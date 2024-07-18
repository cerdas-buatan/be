package module

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
)

func HomeGaysdisal(c *fiber.Ctx, db *mongo.Database) error {
    return c.SendString("Welcome to Gasydisal Bot!")
}

func NotFound(c *fiber.Ctx) error {
    return c.Status(fiber.StatusNotFound).SendString("404 Not Found")
}

