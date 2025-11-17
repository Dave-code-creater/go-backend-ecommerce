package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Database is not connected %v\n", err)
	}
	db.AutoMigrate(&domain.User{})
	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}
	setUpRoutes(rh)
	app.Listen("localhost:9000")
}

func setUpRoutes(rh *rest.RestHandler) {

	// user handler
	handlers.SetupUserRoutes(rh)
	// transactions
	// catalog
}
