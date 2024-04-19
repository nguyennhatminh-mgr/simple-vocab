package routes

import (
	"simple-vocab/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	// Create rdoutes group.
	route := app.Group("/api/v1")

	route.Get("/categories", controllers.GetCategoriesHandler)

	route.Get("/vocabularies", controllers.GetVocabulariesHandler)
	route.Post("/vocabulary", controllers.AddVocabularyHandler)
	route.Get("/vocabulary/:id", controllers.GetVocabularyByIdHandler)
	route.Put("/vocabulary/:i", controllers.UpdatedVocabularyHandler)
	route.Delete("/vocabulary/:id", controllers.DeleteVocabularyHandler)

	route.Get("/users", controllers.GetUsersHandler)
}