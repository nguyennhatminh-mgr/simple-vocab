package main

import (
	"fmt"
	"log"
	"os"
	"simple-vocab/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    app := fiber.New()
	
		errEnv := godotenv.Load()
		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}

		dbUser := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")

		dbUrl := fmt.Sprintf("host=db user=%s password=%s dbname=simple_vocabulary port=5432 sslmode=disable", dbUser, dbPassword)

		db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to the database")
		}

		// Middleware to attach the db to the context
		app.Use(func(c *fiber.Ctx) error {
			c.Locals("db", db)
			return c.Next()
		})

    routes.PrivateRoutes(app)

    app.Listen(":8000")
}
