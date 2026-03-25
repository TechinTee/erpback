package main
import (
	"os"
    "fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
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
    fmt.Println("Sussessfully running on port 3000")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Listen(":" + port)
}