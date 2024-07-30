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

	app.Get("/total_clicks", func(c fiber.Ctx) error { 
		clickData := map[string]interface{}{
			"total_clicks": backend.GetTotalClicks(),
		}
		return c.JSON(clickData)
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
		err := backend.ClickHandler(req.User, req.Timestamp)

		if err != nil {
			log.Printf("Error in ClickHandler: %v", err)
		}

		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString("Received click")

	
	})

	log.Fatal(app.Listen(":3000"))
}

