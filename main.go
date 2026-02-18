package main

import "github.com/gofiber/fiber/v3"
import "fmt"

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World!!!")
    })
    fmt.Println("Sussess")
    app.Listen(":3000")
}