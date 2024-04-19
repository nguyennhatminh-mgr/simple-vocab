package controllers

import (
	"simple-vocab/app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetVocabulariesHandler(c *fiber.Ctx) error {
	var vocabularies []models.Vocabulary

	db := c.Locals("db").(*gorm.DB)

	result := db.Raw(`SELECT * FROM "Vocabularies"`).Scan(&vocabularies)

	if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(vocabularies)
}

func GetVocabularyByIdHandler(c* fiber.Ctx) error {
	id := c.Params("id")

	db := c.Locals("db").(*gorm.DB)

	var currentVocabulary models.Vocabulary

	selectItemQuery := `SELECT * FROM "Vocabularies" WHERE "Vocabularies".id = ?`

	resultCurrentVocabulary := db.Raw(selectItemQuery, id).Scan(&currentVocabulary)

	if resultCurrentVocabulary.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": resultCurrentVocabulary.Error.Error()})
	}

	if (currentVocabulary.Id == 0) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "message": "Not found vocabulary"})
	}

	return c.JSON(currentVocabulary)
}

func AddVocabularyHandler(c* fiber.Ctx) error {
	data := new(models.Vocabulary)
	if err := c.BodyParser(data); err != nil {
		return c.SendString("Some error when parsing")
	}

	var insertedId int

	db := c.Locals("db").(*gorm.DB)

	insertQuery := `INSERT INTO "Vocabularies" (word, meaning, category_id) VALUES (?, ?, ?) RETURNING id`

	result := db.Raw(insertQuery, data.Word, data.Meaning, data.CategoryId).Scan(&insertedId)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	data.Id = insertedId

	return c.JSON(data)
}

func DeleteVocabularyHandler(c* fiber.Ctx) error {
	id := c.Params("id")

	db := c.Locals("db").(*gorm.DB)

	var currentVocabulary models.Vocabulary

	selectItemQuery := `SELECT * FROM "Vocabularies" WHERE "Vocabularies".id = ?`

	resultCurrentVocabulary := db.Raw(selectItemQuery, id).Scan(&currentVocabulary)

	if resultCurrentVocabulary.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": resultCurrentVocabulary.Error.Error()})
	}

	if (currentVocabulary.Id == 0) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "message": "Not found"})
	}

	deleteQuery := `DELETE FROM "Vocabularies" WHERE "Vocabularies".id = ?`

	result := db.Exec(deleteQuery, id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(fiber.Map{ "message": "Deleted successfully"})
}

func UpdatedVocabularyHandler(c* fiber.Ctx) error {
	id := c.Params("id")

	data := new(models.Vocabulary)
	if err := c.BodyParser(data); err != nil {
		return c.SendString("Some error when parsing")
	}

	db := c.Locals("db").(*gorm.DB)

	var currentVocabulary models.Vocabulary

	selectItemQuery := `SELECT * FROM "Vocabularies" WHERE "Vocabularies".id = ?`

	resultCurrentVocabulary := db.Raw(selectItemQuery, id).Scan(&currentVocabulary)

	if resultCurrentVocabulary.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": resultCurrentVocabulary.Error.Error()})
	}

	if (currentVocabulary.Id == 0) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "message": "Not found vocabulary"})
	}

	updateQuery := `UPDATE "Vocabularies" 
			SET word = ?, meaning = ?, category_id = ?
			WHERE "Vocabularies".id = ?`

	result := db.Exec(updateQuery, data.Word, data.Meaning, data.CategoryId, id)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	data.Id, _ = strconv.Atoi(id)

	return c.JSON(fiber.Map{ "message": "Updated successfully", "data": data })
}