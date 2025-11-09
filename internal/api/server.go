package api

import (
	"go-ecommerce-app/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	app.Listen("localhost:9000")
}
