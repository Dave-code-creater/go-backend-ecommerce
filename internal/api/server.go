package api

import (
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app/config"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Listen("localhost:9000")
}
