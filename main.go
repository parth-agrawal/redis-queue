package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/parth-agrawal/redis-queue/cmd/backend"
)


func main () { 
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New())


	app.Get("/", func(c fiber.Ctx) error { 
		return c.SendString("Hello, World!")
	})

	app.Post("/click", func(c fiber.Ctx) error {
		c.Set("Content-Type", "application/json")

		req := struct {
			User      string `json:"user"`
			Timestamp int `json:"timestamp"`
		}{}
	
			// ... existing code ...
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		// ... existing code ...

		// Use the parsed request data instead of FormValue
		backend.ClickHandler(req.User, req.Timestamp)
	
		return c.SendString("Received click")
	})

	log.Fatal(app.Listen(":3000"))
}

