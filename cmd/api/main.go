package api

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Main(){
	fmt.Println("API server is starting...")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods:  "GET, POST, PUT, DELETE, OPTIONS",
	}))
	RegisterRoutes(app)
	log.Fatal(app.Listen(":4000"))
}