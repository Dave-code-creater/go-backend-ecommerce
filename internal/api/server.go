package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	restHandler := &rest.RestHandler{
		App: app,
	}
	setUpRoutes(restHandler)
	app.Listen("localhost:9000")
}

func setUpRoutes(rh *rest.RestHandler) {

	// user handler
	handlers.SetupUserRoutes(rh)
	// transactions
	// catalog
}
