package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error{
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error{
	return c.SendString("single book")
}

func NewBook(c *fiber.Ctx) error{
	return c.SendString("New Book")
}

func DeleteBook(c *fiber.Ctx) error{
	return c.SendString("Delete Book")
}