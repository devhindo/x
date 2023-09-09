package main

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func main() {
	server := fiber.New()
	server.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello, xCLI 👋!")
        return c.SendString("Hello, World 👋!")
    })

    server.Listen(":10000")
}