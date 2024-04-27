package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
}

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

	setUpRoutes(app)
    // Define a route for the GET method on the root path '/'
    // app.Get("/", func(c fiber.Ctx) error {
    //     // Send a string response to the client
    //     return c.SendString("Hello, World 👋!")
    // })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}