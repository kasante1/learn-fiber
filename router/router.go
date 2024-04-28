package router

import (
	"github.com/kasante1/learn-fiber/handler"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {

	api := app.Group("/api/v1")
	v1 := api.Group("book")

	v1.Get("/", handler.GetBooks)
	v1.Get("/:id", handler.GetBook)
	v1.Post("/", handler.CreateBook)
	v1.Put("/:id", handler.UpdateBook)
	v1.Delete("/:id", handler.DeleteBookByID)

}
