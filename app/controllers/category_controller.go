package controllers

import (
	"simple-vocab/app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetCategoriesHandler(c *fiber.Ctx) error {
	var categories []models.Category

	db := c.Locals("db").(*gorm.DB)

	result := db.Raw(`SELECT * FROM "Categories"`).Scan(&categories)

	if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(categories)
}