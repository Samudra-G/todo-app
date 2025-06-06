package api

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

func Main(){
	fmt.Println("API server is starting...")

	app := fiber.New()
	RegisterRoutes(app)
	log.Fatal(app.Listen(":4000"))
}