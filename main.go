package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Printf("Hello, World")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
 
	PORT := os.Getenv("PORT")

	app := fiber.New()

	todos := []Todo{}

	// GET /api/todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos) // SendString() is a method that sends a string to the client
	})
	// POST /api/todos
	app.Post("api/todos", func(c *fiber.Ctx) error {
		todo_create := new(Todo)
		if err := c.BodyParser(todo_create); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if todo_create.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}
		todo_create.ID = len(todos) + 1
		todos = append(todos, *todo_create)
		return c.Status(201).JSON(todo_create)
	})
	//Patch /api/todos/:id
	app.Patch("api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if strconv.Itoa(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})

	})

	//Delete a todo
	app.Delete("api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if strconv.Itoa(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": "Todo deleted"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})

	})

	log.Fatal(app.Listen(":" + PORT))
}
