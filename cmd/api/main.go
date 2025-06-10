package api

import (
	"fmt"
	"log"
	"os"
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
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" 
	}
	log.Fatal(app.Listen(":" + port))
}