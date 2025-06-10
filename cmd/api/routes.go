package api

import (
	_ "github.com/Samudra-G/todo-app/docs"
	"github.com/Samudra-G/todo-app/internal/todo"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/arsmn/fiber-swagger/v2"
)

var todos todo.Todo
const filename = ".todo.json"

func RegisterRoutes(app *fiber.App) {
	todos.Load(filename)

	app.Get("/todos", GetTodos)
	app.Post("/todos", CreateTodo)
	app.Put("/todos/:id/complete", CompleteTodo)
	app.Delete("/todos/:id", DeleteTodo)
	app.Get("/health", HealthCheck)
	app.Get("/docs/*", swagger.HandlerDefault)
}

// GetTodos godoc
// @Summary Get all todos
// @Description Returns all todo items
// @Tags todos
// @Produce json
// @Success 200 {array} todo.Item
// @Router /todos [get]
func GetTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Adds a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body map[string]string true "Todo task"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /todos [post]
func CreateTodo(c *fiber.Ctx) error {
	var body struct {
		Task string `json:"task"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if body.Task == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Task cannot be empty"})
	}
	if err := todos.Load(filename); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to load todos"})
	}
	todos.Add(body.Task)
	todos.Store(filename)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Todo item added successfully"})
}

// CompleteTodo godoc
// @Summary Mark todo as complete
// @Description Marks the given todo item as complete
// @Tags todos
// @Param id path int true "Todo ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id}/complete [put]
func CompleteTodo(c *fiber.Ctx) error {
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
	return c.JSON(fiber.Map{"message": "Marked as complete"})
}

// DeleteTodo godoc
// @Summary Delete a todo item
// @Description Deletes a todo item by ID
// @Tags todos
// @Param id path int true "Todo ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [delete]
func DeleteTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid index")
	}
	if err := todos.Delete(id); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}
	todos.Store(filename)
	return c.JSON(fiber.Map{"message": "Todo item deleted successfully"})
}

// HealthCheck godoc
// @Summary Health check
// @Description Check if server is running
// @Tags system
// @Success 200
// @Router /health [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
