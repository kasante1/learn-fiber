package main

import (
	// "fmt"
	// "log"
	"github.com/kasante1/learn-fiber/database"
	"github.com/kasante1/learn-fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"

)


func main() {
    // Initialize a new Fiber app
	database.Connect()

    app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error{
		return c.SendStatus(404) // 404 not found
	})

	app.Listen(":8080")


}