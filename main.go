package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, ensure DATABASE_URL is set in environment")
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/commingsoon", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"textname": "COMMING SOON",
			},
			"message": "Fetch data successfully",
		})
	})
	app.Get("/ping-db", func(c fiber.Ctx) error {
		databaseURL := os.Getenv("DATABASE_URL")
		if databaseURL == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "DATABASE_URL is not set",
			})
		}

		db, err := sql.Open("postgres", databaseURL)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to open DB: " + err.Error(),
			})
		}
		defer db.Close()

		if err := db.Ping(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "DB Ping failed: " + err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Connected to Neon DB!",
		})
	})

	fmt.Println("Sussessfully running on port 3000")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}
