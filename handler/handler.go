package handler

import (
	"github.com/kasante1/learn-fiber/database"
	"github.com/kasante1/learn-fiber/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateBook(c *fiber.Ctx) error {
	db := database.DB.Db
	book := new(book.Book)

	err := c.BodyParser(book)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "invalid iput", "data": err})
	}

	err = db.Create(&book).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "could not create user", "data": err})
	}

	// return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "book created successfully", "data": book})
}

// get all books

func GetBooks(c *fiber.Ctx) error {
	db := database.DB.Db
	var books []book.Book

	db.Find(&books)

	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "books not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "books found", "data": books})
}

func GetBook(c *fiber.Ctx) error {
	db := database.DB.Db
	
	// get id params
	id := c.Params("id")

	var book book.Book

	// find single user in the db by id
	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "book not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "book found", "data": book})
}


func UpdateBook(c *fiber.Ctx) error {
	type updateBook struct {
		Title string `json:"title"`
	}

	db := database.DB.Db

	var book book.Book

	// get id params
	id := c.Params("id")

	// find single book in the db by id
	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "book not found", "data": nil})

	}

	var updateBookData updateBook
	err := c.BodyParser(&updateBookData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "invalid input",
			"data": err,
		})
	}
	book.Title = updateBookData.Title

	// save the changes
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": " book found",
		"data": book,
	})
}

func DeleteBookByID(c *fiber.Ctx) error {
	db := database.DB.Db

	var book book.Book

	//get id params
	id := c.Params("id")

	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error", 
			"message": "book not found",
			"data": nil,
		})
	}

	err := db.Delete(&book, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": " delete book failed",
			"data": nil,
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"message": "delete book success",
	})
}