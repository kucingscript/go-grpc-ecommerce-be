package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/config"
	"github.com/kucingscript/go-grpc-ecommerce-be/internal/handler/product"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panicf("failed to load config: %v", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS_ALLOWED_ORIGINS,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
	}))

	app.Get("/storage/product/:filename", product.GetProductImageHandler)
	app.Post("/product/upload", product.UploadProductImageHandler)

	app.Listen(cfg.REST_PORT)
}
