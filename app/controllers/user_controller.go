package controllers

import (
	"simple-vocab/app/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func GetUsersHandler(c *fiber.Ctx) error {
	var users []models.User

	db := c.Locals("db").(*gorm.DB)

	result := db.Raw(`SELECT * FROM "Users"`).Scan(&users)

	if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(users)
}