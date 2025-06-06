package api

import (
	"github.com/Samudra-G/todo-app/internal/todo"
	"github.com/gofiber/fiber/v2"
)

var todos todo.Todo
const filename = ".todo.json"

func RegisterRoutes(app *fiber.App){
	todos.Load(filename)

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Post("/todos", func(c *fiber.Ctx) error{
		var body struct{
			Task string `json:"task"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		if body.Task == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Task cannot be empty",
			})
		}
		if err := todos.Load(filename); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to load todos"})
		}

		todos.Add(body.Task)
		todos.Store(filename)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Todo item added successfully",
		})
	})

	app.Put("/todos/:id/complete", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil || id <= 0 {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid index")
		}
		if err := todos.Complete(id); err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
		}
		if err := todos.Store(filename); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save changes"})
		}	
		todos.Store(filename)
		return c.JSON(fiber.Map{"message": "Marked as complete"})
	})

	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil || id <= 0{
			return fiber.NewError(fiber.StatusBadRequest, "Invalid index")
		}
		if err := todos.Delete(id); err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
		}
		todos.Store(filename)
		return c.JSON(fiber.Map{"message": "Todo item deleted successfully"})
	})
}